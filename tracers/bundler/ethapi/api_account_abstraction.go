package ethapi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type parsedTransaction struct {
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       string         `json:"blockNumber"`
	From              common.Address `json:"from"`
	CumulativeGasUsed string         `json:"cumulativeGasUsed"`
	GasUsed           string         `json:"gasUsed"`
	Logs              []*types.Log   `json:"logs"`
	LogsBloom         types.Bloom    `json:"logsBloom"`
	TransactionHash   common.Hash    `json:"transactionHash"`
	TransactionIndex  string         `json:"transactionIndex"`
	EffectiveGasPrice string         `json:"effectiveGasPrice"`
}

// UserOperation is a struct that represents a user operation at ERC-4337.
type UserOperation struct {
	Sender               common.Address `json:"sender"`
	Nonce                *big.Int       `json:"nonce"`
	InitCode             []byte         `json:"initCode"`
	CallData             []byte         `json:"callData"`
	CallGasLimit         *big.Int       `json:"callGasLimit"`
	VerificationGasLimit *big.Int       `json:"verificationGasLimit"`
	PreVerificationGas   *big.Int       `json:"preVerificationGas"`
	MaxFeePerGas         *big.Int       `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int       `json:"maxPriorityFeePerGas"`
	PaymasterAndData     []byte         `json:"paymasterAndData"`
	Signature            []byte         `json:"signature"`
}

// UserOperationReceipt is a struct that represents a user operation receipt in ERC-4337.
type UserOperationReceipt struct {
	UserOpHash    common.Hash        `json:"userOpHash"`
	Sender        common.Address     `json:"sender"`
	Paymaster     common.Address     `json:"paymaster"`
	Nonce         string             `json:"nonce"`
	Success       bool               `json:"success"`
	ActualGasCost string             `json:"actualGasCost"`
	ActualGasUsed string             `json:"actualGasUsed"`
	From          common.Address     `json:"from"`
	Receipt       *parsedTransaction `json:"receipt"`
	Logs          []*types.Log       `json:"logs"`
}

func (s *TransactionAPI) SendUserOperation(ctx context.Context, args types.Transactions, creationBlock *big.Int) (common.Hash, error) {
	log.Warn("SendUserOperation is not implemented")
	return common.Hash{}, nil
}

func (s *TransactionAPI) GetUserOperationReceipt(ctx context.Context, hash common.Hash) (*UserOperationReceipt, error) {
	log.Warn("GetUserOperationReceipt is not implemented")
	return nil, nil
}

func (s *TransactionAPI) SupportedEntrypoints() ([]string, error) {
	log.Warn("SupportedEntrypoints is not implemented")
	return nil, nil
}

func (s *TransactionAPI) GetUserOperationByHash(ctx context.Context, hash common.Hash) (*UserOperationReceipt, error) {
	log.Warn("GetUserOperationByHash is not implemented")
	return nil, nil
}

// GasEstimates provides estimate values for all gas fields in a UserOperation.
type GasEstimates struct {
	PreVerificationGas   *big.Int `json:"preVerificationGas"`
	VerificationGasLimit *big.Int `json:"verificationGasLimit"`
	CallGasLimit         *big.Int `json:"callGasLimit"`

	// TODO: Deprecate in v0.7
	VerificationGas *big.Int `json:"verificationGas"`
}

func (api *BlockChainAPI) EstimateUserOperationGas(ctx context.Context, args types.Transactions, creationBlock *big.Int) (*GasEstimates, error) {
	log.Warn("EstimateUserOperationGas is not implemented")
	return nil, nil
}
