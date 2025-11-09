package across_spoke

import "math/big"

// FundsDepositedEvent 存储 Across FundsDeposited 事件的所有数据
type FundsDepositedEvent struct {
	// Indexed 字段（在 Topics 中）
	DestinationChainId *big.Int `json:"destination_chain_id"` // Topics[1]
	DepositId          *big.Int `json:"deposit_id"`           // Topics[2]
	Depositor          string   `json:"depositor"`            // Topics[3] - bytes32 转换为 address

	// Non-indexed 字段（在 Data 中）
	InputToken          string   `json:"input_token"`  // bytes32 转换为 address
	OutputToken         string   `json:"output_token"` // bytes32 转换为 address
	InputAmount         *big.Int `json:"input_amount"`
	OutputAmount        *big.Int `json:"output_amount"`
	QuoteTimestamp      uint32   `json:"quote_timestamp"`
	FillDeadline        uint32   `json:"fill_deadline"`
	ExclusivityDeadline uint32   `json:"exclusivity_deadline"`
	Recipient           string   `json:"recipient"`         // bytes32 转换为 address
	ExclusiveRelayer    string   `json:"exclusive_relayer"` // bytes32 转换为 address
	Message             []byte   `json:"message"`
}
