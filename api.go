package blockchain

import (
	"fmt"
)

// BitgesellBlockchainSDK represents the Bitgesell Blockchain SDK
type BitgesellBlockchainSDK struct {
	Blockchain Blockchain
	Tx         Transaction
	Mempool    Mempool
	Address    Address
}

// NewBitgesellBlockchainSDK creates a new instance of the Bitgesell Blockchain SDK
func NewBitgesellBlockchainSDK(config SDKConfig) *BitgesellBlockchainSDK {
	return &BitgesellBlockchainSDK{
		Blockchain: Blockchain{apiV1URL: config.BaseAPIURL},
		Tx:         Transaction{apiV1URL: config.BaseAPIURL},
		Mempool:    Mempool{apiV1URL: config.BaseAPIURL},
		Address:    Address{apiV1URL: config.BaseAPIURL},
	}
}
