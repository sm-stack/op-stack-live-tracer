package ethapi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"sync"
)

// AddrLocker is a mutex locker for addresses, declared to prevent build errors.
type AddrLocker struct {
	mu    sync.Mutex
	locks map[common.Address]*sync.Mutex
}

// TransactionAPI is a struct that contains a backend, an address locker, and a signer.
// Declared to prevent build errors.
type TransactionAPI struct {
	b         eth.Backend
	nonceLock *AddrLocker
	signer    types.Signer
}

// BlockChainAPI is a struct that contains a backend. Declared to prevent build errors.
type BlockChainAPI struct {
	b eth.Backend
}
