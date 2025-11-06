package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// This tool updates min_amount and max_amount for ETH, USDC, USDT, USD1 in t_across_routes.
// Defaults:
//   - ETH_MAX=10, ETH_MIN=0
//   - USDC_MAX=10000, USDC_MIN=0
//   - USDT_MAX=10000, USDT_MIN=0
//   - USD1_MAX=10000, USD1_MIN=0
//
// Override via env vars as needed.
// DB connection:
//   - Use ACROSS_TEST_DSN for full DSN, or ACROSS_DB_NAME for db name with default host/user/pass.
func main() {
	dsn := os.Getenv("ACROSS_TEST_DSN")
	if dsn == "" {
		host := "test.cjqgi44oi65t.us-east-1.rds.amazonaws.com"
		user := "admin"
		pass := "blJDcXp6MZe3rCCrjLgR"
		dbName := os.Getenv("ACROSS_DB_NAME")
		if dbName == "" {
			dbName = "test"
		}
		dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&charset=utf8mb4&loc=Local", user, pass, host, dbName)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("ping db error: %v", err)
	}

	// Read amounts from env (use strings to avoid precision issues)
	// Core chain minima (for chains: Ethereum, Arbitrum, Base, Optimism, Linea)
	ethCoreMin := getenvDefault("ETH_CORE_MIN", getenvDefault("ETH_MIN", "1"))
	// Non-core chain minima: 非 Ethereum/Arbitrum/Base/Optimism/Linea 链，ETH 最小值 0.1
	ethNonCoreMin := getenvDefault("ETH_NONCORE_MIN", "0.1")
	ethMax := getenvDefault("ETH_MAX", "10")

	// Stablecoin minima and maxima
	stableCoreMin := getenvDefault("STABLE_CORE_MIN", getenvDefault("STABLE_MIN", getenvDefault("USDC_MIN", "500")))
	// 非 Ethereum/Arbitrum/Base/Optimism/Linea/Solana/BNBChain 链，稳定币最小值 100
	stableNonCoreMin := getenvDefault("STABLE_NONCORE_MIN", "100")
	usdcMax := getenvDefault("USDC_MAX", "10000")
	usdtMax := getenvDefault("USDT_MAX", "10000")
	usd1Max := getenvDefault("USD1_MAX", "10000")

	// Core chain ID sets (逗号分隔). 默认使用常见 EVM ChainId，Solana/BNB 可按需覆盖。
	// 你可以通过环境变量覆盖：CORE_CHAINS_ETH, CORE_CHAINS_STABLE
	ethCoreChains := getenvDefault("CORE_CHAINS_ETH", "1,42161,8453,10,59144")
	stableCoreChains := getenvDefault("CORE_CHAINS_STABLE", "1,42161,8453,10,59144,56,34268394551451")

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("begin tx error: %v", err)
	}
	totalAffected := int64(0)

	// 1) 全局设置最大值（不区分链）
	{
		maxStmt, err := tx.Prepare(`UPDATE t_across_routes SET max_amount = ?, updated_at = NOW() WHERE origin_token_symbol = ? OR destination_token_symbol = ?`)
		if err != nil {
			_ = tx.Rollback()
			log.Fatalf("prepare max update stmt error: %v", err)
		}
		defer maxStmt.Close()

		// ETH
		if res, err := maxStmt.Exec(ethMax, "ETH", "ETH"); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update ETH max error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated ETH max rows affected: %d (max=%s)", n, ethMax)
		}
		// USDC
		if res, err := maxStmt.Exec(usdcMax, "USDC", "USDC"); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update USDC max error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated USDC max rows affected: %d (max=%s)", n, usdcMax)
		}
		// USDT
		if res, err := maxStmt.Exec(usdtMax, "USDT", "USDT"); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update USDT max error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated USDT max rows affected: %d (max=%s)", n, usdtMax)
		}
		// USD1
		if res, err := maxStmt.Exec(usd1Max, "USD1", "USD1"); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update USD1 max error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated USD1 max rows affected: %d (max=%s)", n, usd1Max)
		}
	}

	// 构造 IN 子句字符串（确保只含数字和逗号）
	inEth := normalizeInClause(ethCoreChains)
	inStable := normalizeInClause(stableCoreChains)
	if inEth == "" || inStable == "" {
		_ = tx.Rollback()
		log.Fatalf("core chain ID sets are empty; please set CORE_CHAINS_ETH/CORE_CHAINS_STABLE")
	}

	// 2) 设置核心链最小值（两端均在核心链集合中）
	{
		// ETH 核心链最小值
		query := fmt.Sprintf(`UPDATE t_across_routes
            SET min_amount = ?, updated_at = NOW()
            WHERE (origin_token_symbol = 'ETH' OR destination_token_symbol = 'ETH')
            AND (origin_chain_id IN (%s) AND destination_chain_id IN (%s))`, inEth, inEth)
		if res, err := tx.Exec(query, ethCoreMin); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update ETH core min error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated ETH core min rows affected: %d (min=%s)", n, ethCoreMin)
		}

		// 稳定币（USDC/USDT/USD1）核心链最小值
		queryStable := fmt.Sprintf(`UPDATE t_across_routes
            SET min_amount = ?, updated_at = NOW()
            WHERE (origin_token_symbol IN ('USDC','USDT','USD1') OR destination_token_symbol IN ('USDC','USDT','USD1'))
            AND (origin_chain_id IN (%s) AND destination_chain_id IN (%s))`, inStable, inStable)
		if res, err := tx.Exec(queryStable, stableCoreMin); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update stable core min error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated stable core min rows affected: %d (min=%s)", n, stableCoreMin)
		}
	}

	// 3) 设置非核心链最小值（任一端不在核心链集合中）
	{
		// ETH 非核心链最小值 0.1
		query := fmt.Sprintf(`UPDATE t_across_routes
            SET min_amount = ?, updated_at = NOW()
            WHERE (origin_token_symbol = 'ETH' OR destination_token_symbol = 'ETH')
            AND (origin_chain_id NOT IN (%s) OR destination_chain_id NOT IN (%s))`, inEth, inEth)
		if res, err := tx.Exec(query, ethNonCoreMin); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update ETH non-core min error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated ETH non-core min rows affected: %d (min=%s)", n, ethNonCoreMin)
		}

		// 稳定币（USDC/USDT/USD1）非核心链最小值 100
		queryStable := fmt.Sprintf(`UPDATE t_across_routes
            SET min_amount = ?, updated_at = NOW()
            WHERE (origin_token_symbol IN ('USDC','USDT','USD1') OR destination_token_symbol IN ('USDC','USDT','USD1'))
            AND (origin_chain_id NOT IN (%s) OR destination_chain_id NOT IN (%s))`, inStable, inStable)
		if res, err := tx.Exec(queryStable, stableNonCoreMin); err != nil {
			_ = tx.Rollback()
			log.Fatalf("update stable non-core min error: %v", err)
		} else if n, _ := res.RowsAffected(); n >= 0 {
			totalAffected += n
			log.Printf("Updated stable non-core min rows affected: %d (min=%s)", n, stableNonCoreMin)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("commit error: %v", err)
	}
	log.Printf("All updates committed. Total rows affected: %d", totalAffected)
}

func getenvDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// normalizeInClause ensures the IN clause contains only integers separated by commas
// e.g. "1, 10,42161" -> "1,10,42161"; invalid items are discarded.
func normalizeInClause(csv string) string {
	parts := strings.Split(csv, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		// allow only digits
		valid := true
		for i := 0; i < len(p); i++ {
			if p[i] < '0' || p[i] > '9' {
				valid = false
				break
			}
		}
		if valid {
			out = append(out, p)
		}
	}
	return strings.Join(out, ",")
}
