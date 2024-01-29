package main

import (
	"fmt"
	"github.com/naftalimurgor/go-bitgesell-toolkit"
)

func main() {
	// Example usage of the Bitgesell Blockchain SDK
	config := blockchain.SDKConfig.SDKConfig{BaseAPIURL: "https://api.bitaps.com/bgl/v1/blockchain"}
	bitgesellSDK := blockchain.bitgesell.NewBitgesellBlockchainSDK(config)

	// Example: Access Blockchain SDK methods
	block, err := blockchain.Blockchain.GetBlockByHash("your_block_hash")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Block:", block)

	// Example: Access Transaction SDK methods
	tx, err := bitgesellSDK.Tx.GetTransactionByHash("your_transaction_hash")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Transaction:", tx)

	// Example: Access Mempool SDK methods
	mempoolTransactions, err := bitgesellSDK.Mempool.GetMempoolTransactions(10, "asc", 0, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Mempool Transactions:", mempoolTransactions)
}