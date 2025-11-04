package loader

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/owlto-dao/utils-go/alert"
)

type AcrossRoute struct {
    ID                     int64  `json:"id"`
    OriginChainId          int64  `json:"originChainId"`
    OriginToken            string `json:"originToken"`
    DestinationChainId     int64  `json:"destinationChainId"`
    DestinationToken       string `json:"destinationToken"`
    OriginTokenSymbol      string `json:"originTokenSymbol"`
    DestinationTokenSymbol string `json:"destinationTokenSymbol"`
    IsNative               bool   `json:"isNative"`
    MinAmount              string `json:"minAmount"`
    MaxAmount              string `json:"maxAmount"`
    FeeRate                string `json:"feeRate"`
    IsActive               bool   `json:"isActive"`
}

type AcrossManager struct {
    originToDestRoutes     map[int64]map[int64][]*AcrossRoute
    tokenPairRoutes        map[string]map[string][]*AcrossRoute
    chainTokenRoutes       map[int64]map[string][]*AcrossRoute
    symbolPairRoutes       map[string]map[string][]*AcrossRoute
    routeById              map[int64]*AcrossRoute
    db                     *sql.DB
    alerter                alert.Alerter
    mutex                  *sync.RWMutex
    defaultFeeRate         string
    routeRefreshInterval   time.Duration
    stopCh                 chan struct{}
}

func NewAcrossManager(db *sql.DB, alerter alert.Alerter) *AcrossManager {
    return &AcrossManager{
        originToDestRoutes:   make(map[int64]map[int64][]*AcrossRoute),
        tokenPairRoutes:      make(map[string]map[string][]*AcrossRoute),
        chainTokenRoutes:     make(map[int64]map[string][]*AcrossRoute),
        symbolPairRoutes:     make(map[string]map[string][]*AcrossRoute),
        routeById:            make(map[int64]*AcrossRoute),
        db:                   db,
        alerter:              alerter,
        mutex:                new(sync.RWMutex),
        defaultFeeRate:       "0",
        routeRefreshInterval: 24 * time.Hour,
    }
}

func (mgr *AcrossManager) FetchRoutesFromAPI() ([]*AcrossRoute, error) {
	resp, err := http.Get("https://app.across.to/api/available-routes")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch routes from API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch routes from API: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var routes []*AcrossRoute
	if err := json.Unmarshal(body, &routes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal routes: %w", err)
	}

	return routes, nil
}

func (mgr *AcrossManager) SyncRoutesWithDB() error {
    routes, err := mgr.FetchRoutesFromAPI()
    if err != nil {
        return fmt.Errorf("failed to fetch routes for DB sync: %w", err)
    }

    tx, err := mgr.db.Begin()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }

    const chunkSize = 200
    for i := 0; i < len(routes); i += chunkSize {
        end := i + chunkSize
        if end > len(routes) {
            end = len(routes)
        }
        chunk := routes[i:end]

        var sb strings.Builder
        sb.WriteString(`INSERT INTO t_across_routes (
                origin_chain_id, origin_token, destination_chain_id, destination_token,
                origin_token_symbol, destination_token_symbol, is_native,
                min_amount, max_amount, fee_rate, is_active,
                created_at, updated_at
            ) VALUES `)

        args := make([]any, 0, len(chunk)*11)
        for idx, route := range chunk {
            if idx > 0 {
                sb.WriteString(",")
            }
            sb.WriteString("(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
            args = append(args,
                route.OriginChainId, route.OriginToken, route.DestinationChainId, route.DestinationToken,
                route.OriginTokenSymbol, route.DestinationTokenSymbol, route.IsNative,
                route.MinAmount, route.MaxAmount, route.FeeRate, route.IsActive,
            )
        }

        sb.WriteString(` ON DUPLICATE KEY UPDATE
                origin_token_symbol = VALUES(origin_token_symbol),
                destination_token_symbol = VALUES(destination_token_symbol),
                is_native = VALUES(is_native),
                min_amount = VALUES(min_amount),
                max_amount = VALUES(max_amount),
                fee_rate = VALUES(fee_rate),
                is_active = VALUES(is_active),
                updated_at = NOW();`)

        if _, err := tx.Exec(sb.String(), args...); err != nil {
            tx.Rollback()
            return fmt.Errorf("failed to upsert routes chunk [%d:%d]: %w", i, end, err)
        }
    }

    return tx.Commit()
}

// LoadAllRoutes queries active routes from DB and rebuilds in-memory indexes atomically.
func (mgr *AcrossManager) LoadAllRoutes() error {
    rows, err := mgr.db.Query(`SELECT id, origin_chain_id, origin_token, destination_chain_id, destination_token,
        origin_token_symbol, destination_token_symbol, is_native, min_amount, max_amount, fee_rate, is_active
        FROM t_across_routes`)
    if err != nil {
        return fmt.Errorf("query t_across_routes error: %w", err)
    }
    if rows == nil {
        return fmt.Errorf("query t_across_routes returned nil rows")
    }
    defer rows.Close()

    newOriginToDest := make(map[int64]map[int64][]*AcrossRoute)
    newTokenPair := make(map[string]map[string][]*AcrossRoute)
    newChainToken := make(map[int64]map[string][]*AcrossRoute)
    newSymbolPair := make(map[string]map[string][]*AcrossRoute)
    newByID := make(map[int64]*AcrossRoute)

    for rows.Next() {
        var r AcrossRoute
        if err := rows.Scan(
            &r.ID,
            &r.OriginChainId,
            &r.OriginToken,
            &r.DestinationChainId,
            &r.DestinationToken,
            &r.OriginTokenSymbol,
            &r.DestinationTokenSymbol,
            &r.IsNative,
            &r.MinAmount,
            &r.MaxAmount,
            &r.FeeRate,
            &r.IsActive,
        ); err != nil {
            mgr.alerter.AlertText("scan t_across_routes row error", err)
            continue
        }

        // Normalize strings for indexing keys
        oTok := strings.ToLower(strings.TrimSpace(r.OriginToken))
        dTok := strings.ToLower(strings.TrimSpace(r.DestinationToken))
        oSym := strings.ToLower(strings.TrimSpace(r.OriginTokenSymbol))
        dSym := strings.ToLower(strings.TrimSpace(r.DestinationTokenSymbol))

        // origin -> dest index
        if _, ok := newOriginToDest[r.OriginChainId]; !ok {
            newOriginToDest[r.OriginChainId] = make(map[int64][]*AcrossRoute)
        }
        newOriginToDest[r.OriginChainId][r.DestinationChainId] = append(newOriginToDest[r.OriginChainId][r.DestinationChainId], &r)

        // token pair index
        if _, ok := newTokenPair[oTok]; !ok {
            newTokenPair[oTok] = make(map[string][]*AcrossRoute)
        }
        newTokenPair[oTok][dTok] = append(newTokenPair[oTok][dTok], &r)

        // chain token index
        if _, ok := newChainToken[r.OriginChainId]; !ok {
            newChainToken[r.OriginChainId] = make(map[string][]*AcrossRoute)
        }
        newChainToken[r.OriginChainId][oTok] = append(newChainToken[r.OriginChainId][oTok], &r)

        // symbol pair index
        if _, ok := newSymbolPair[oSym]; !ok {
            newSymbolPair[oSym] = make(map[string][]*AcrossRoute)
        }
        newSymbolPair[oSym][dSym] = append(newSymbolPair[oSym][dSym], &r)

        // id index
        newByID[r.ID] = &r
    }
    if err := rows.Err(); err != nil {
        mgr.alerter.AlertText("iterate t_across_routes error", err)
    }

    // Swap caches atomically
    mgr.mutex.Lock()
    mgr.originToDestRoutes = newOriginToDest
    mgr.tokenPairRoutes = newTokenPair
    mgr.chainTokenRoutes = newChainToken
    mgr.symbolPairRoutes = newSymbolPair
    mgr.routeById = newByID
    mgr.mutex.Unlock()

    return nil
}

// LoadRoutes is an alias for LoadAllRoutes for compatibility.
func (mgr *AcrossManager) LoadRoutes() error { return mgr.LoadAllRoutes() }

// Query by chain pair
func (mgr *AcrossManager) GetRoutesByChains(originChainId, destChainId int64) ([]*AcrossRoute, error) {
    mgr.mutex.RLock()
    destMap, ok := mgr.originToDestRoutes[originChainId]
    if !ok {
        mgr.mutex.RUnlock()
        return nil, fmt.Errorf("no routes for originChainId %d", originChainId)
    }
    routes, ok := destMap[destChainId]
    mgr.mutex.RUnlock()
    if !ok || len(routes) == 0 {
        return nil, fmt.Errorf("no routes for destChainId %d", destChainId)
    }
    return routes, nil
}

// Compatibility alias
func (mgr *AcrossManager) GetRoutesByOriginAndDest(originChainId, destChainId int64) ([]*AcrossRoute, error) {
    return mgr.GetRoutesByChains(originChainId, destChainId)
}

// Query by token pair (addresses)
func (mgr *AcrossManager) GetRoutesByTokenPair(originToken, destToken string) ([]*AcrossRoute, error) {
    oTok := strings.ToLower(strings.TrimSpace(originToken))
    dTok := strings.ToLower(strings.TrimSpace(destToken))
    mgr.mutex.RLock()
    destMap, ok := mgr.tokenPairRoutes[oTok]
    if !ok {
        mgr.mutex.RUnlock()
        return nil, fmt.Errorf("no routes for originToken %s", originToken)
    }
    routes, ok := destMap[dTok]
    mgr.mutex.RUnlock()
    if !ok || len(routes) == 0 {
        return nil, fmt.Errorf("no routes for destToken %s", destToken)
    }
    return routes, nil
}

// Query by chain and a token address (origin side)
func (mgr *AcrossManager) GetRoutesByChainAndToken(chainId int64, tokenAddr string) ([]*AcrossRoute, error) {
    t := strings.ToLower(strings.TrimSpace(tokenAddr))
    mgr.mutex.RLock()
    tokMap, ok := mgr.chainTokenRoutes[chainId]
    if !ok {
        mgr.mutex.RUnlock()
        return nil, fmt.Errorf("no routes for chainId %d", chainId)
    }
    routes, ok := tokMap[t]
    mgr.mutex.RUnlock()
    if !ok || len(routes) == 0 {
        return nil, fmt.Errorf("no routes for token %s on chain %d", tokenAddr, chainId)
    }
    return routes, nil
}

// Query by symbol pair
func (mgr *AcrossManager) GetRoutesBySymbolPair(originSymbol, destSymbol string) ([]*AcrossRoute, error) {
    o := strings.ToLower(strings.TrimSpace(originSymbol))
    d := strings.ToLower(strings.TrimSpace(destSymbol))
    mgr.mutex.RLock()
    destMap, ok := mgr.symbolPairRoutes[o]
    if !ok {
        mgr.mutex.RUnlock()
        return nil, fmt.Errorf("no routes for originSymbol %s", originSymbol)
    }
    routes, ok := destMap[d]
    mgr.mutex.RUnlock()
    if !ok || len(routes) == 0 {
        return nil, fmt.Errorf("no routes for destSymbol %s", destSymbol)
    }
    return routes, nil
}

// Get route by ID
func (mgr *AcrossManager) GetRouteByID(id int64) (*AcrossRoute, bool) {
    mgr.mutex.RLock()
    r, ok := mgr.routeById[id]
    mgr.mutex.RUnlock()
    return r, ok
}

// Validate route supports amount in [min, max]
func (mgr *AcrossManager) ValidateRoute(route *AcrossRoute, amount *big.Int) error {
    if route == nil {
        return fmt.Errorf("nil route")
    }
    if amount == nil || amount.Sign() <= 0 {
        return fmt.Errorf("invalid amount")
    }
    min := new(big.Int)
    max := new(big.Int)
    if _, ok := min.SetString(strings.TrimSpace(route.MinAmount), 10); !ok {
        return fmt.Errorf("invalid minAmount: %s", route.MinAmount)
    }
    if _, ok := max.SetString(strings.TrimSpace(route.MaxAmount), 10); !ok {
        return fmt.Errorf("invalid maxAmount: %s", route.MaxAmount)
    }
    if amount.Cmp(min) < 0 {
        return fmt.Errorf("amount below min")
    }
    if max.Sign() > 0 && amount.Cmp(max) > 0 { // max=0 means unlimited
        return fmt.Errorf("amount above max")
    }
    return nil
}

// InvalidateCache clears in-memory indexes (does not touch DB)
func (mgr *AcrossManager) InvalidateCache() {
    mgr.mutex.Lock()
    mgr.originToDestRoutes = make(map[int64]map[int64][]*AcrossRoute)
    mgr.tokenPairRoutes = make(map[string]map[string][]*AcrossRoute)
    mgr.chainTokenRoutes = make(map[int64]map[string][]*AcrossRoute)
    mgr.symbolPairRoutes = make(map[string]map[string][]*AcrossRoute)
    mgr.routeById = make(map[int64]*AcrossRoute)
    mgr.mutex.Unlock()
}

// Start launches periodic sync using given interval.
func (mgr *AcrossManager) StartPeriodicSync(interval time.Duration) {
    mgr.mutex.Lock()
    if mgr.stopCh != nil {
        mgr.mutex.Unlock()
        return
    }
    mgr.stopCh = make(chan struct{})
    mgr.mutex.Unlock()

    go func() {
        ticker := time.NewTicker(interval)
        defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                if err := mgr.SyncRoutesWithDB(); err != nil {
                    mgr.alerter.AlertText("periodic SyncRoutesWithDB error", err)
                    continue
                }
                if err := mgr.LoadAllRoutes(); err != nil {
                    mgr.alerter.AlertText("periodic LoadAllRoutes error", err)
                }
            case <-mgr.stopCh:
                return
            }
        }
    }()
}

// StopPeriodicSync stops the periodic sync goroutine.
func (mgr *AcrossManager) StopPeriodicSync() {
    mgr.mutex.Lock()
    if mgr.stopCh != nil {
        close(mgr.stopCh)
        mgr.stopCh = nil
    }
    mgr.mutex.Unlock()
}

// Start uses manager's default interval.
func (mgr *AcrossManager) Start() { mgr.StartPeriodicSync(mgr.routeRefreshInterval) }

// GetSupportedChains returns all chain IDs observed in routes (origin and destination).
func (mgr *AcrossManager) GetSupportedChains() []int64 {
    mgr.mutex.RLock()
    chains := make(map[int64]struct{})
    for o, destMap := range mgr.originToDestRoutes {
        chains[o] = struct{}{}
        for d := range destMap {
            chains[d] = struct{}{}
        }
    }
    mgr.mutex.RUnlock()
    res := make([]int64, 0, len(chains))
    for c := range chains {
        res = append(res, c)
    }
    return res
}

// GetSupportedTokens returns token addresses supported on a chain (origin side index).
func (mgr *AcrossManager) GetSupportedTokens(chainId int64) []string {
    mgr.mutex.RLock()
    tokMap, ok := mgr.chainTokenRoutes[chainId]
    if !ok {
        mgr.mutex.RUnlock()
        return nil
    }
    res := make([]string, 0, len(tokMap))
    for t := range tokMap {
        res = append(res, t)
    }
    mgr.mutex.RUnlock()
    return res
}