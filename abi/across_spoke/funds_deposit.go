package across_spoke

import "math/big"

// FundsDepositedEvent 存储 Across FundsDeposited 事件的所有数据
type FundsDepositedEvent struct {
	// Indexed 字段（在 Topics 中）
	DestinationChainId *big.Int // Topics[1]
	DepositId          *big.Int // Topics[2]
	Depositor          string   // Topics[3] - bytes32 转换为 address

	// Non-indexed 字段（在 Data 中）
	InputToken          string // bytes32 转换为 address
	OutputToken         string // bytes32 转换为 address
	InputAmount         *big.Int
	OutputAmount        *big.Int
	QuoteTimestamp      uint32
	FillDeadline        uint32
	ExclusivityDeadline uint32
	Recipient           string // bytes32 转换为 address
	ExclusiveRelayer    string // bytes32 转换为 address
	Message             []byte
}
