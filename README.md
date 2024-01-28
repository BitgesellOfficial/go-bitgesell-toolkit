# go-bitgesell-toolkit
<p align="center">
<img src="doc/img/Bitgesell.png" style="height: 60px;"/>
<p align="center">Sign Bitgesell Transactions offline</p>
</p>

`go-bitgesell-toolkit` is a toolkit for interacting with the Bitgesell blockchain. It provides convenient methods for accessing blockchain, transaction, mempool, and address-related data.

Please see full [API documentation](https://github.com/bitaps-com/bglapiserver/tree/master/api)

## Installation

To use the toolkit, you can install it using the following `go get` command:

```bash
go get github.com/naftalimurgor/go-bitgesell-toolkit
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/naftalimurgor/go-bitgesell-toolkit"
)

func main() {
	// Example usage of the Bitgesell Blockchain SDK
	config := bitgesell.SDKConfig{BaseAPIURL: "https://api.bitaps.com/bgl/v1/blockchain"}
	bitgesellSDK := bitgesell.NewBitgesellBlockchainSDK(config)

	// Example: Access Blockchain SDK methods
	block, err := bitgesellSDK.Blockchain.GetBlockByHash("your_block_hash")
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

	// Example: Access Address SDK methods
	addressTransaction, err := bitgesellSDK.Address.GetAddressTransaction("your_address")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Address Transaction:", addressTransaction)
}
```

## Documentation

For detailed information on available methods and types, refer to the [GoDoc documentation](https://github.com/naftalimurgor/go-bitgesell-toolkit).

## License

This toolkit is open-source and available under the [MIT License](LICENSE). Feel free to contribute, report issues, or use it in your projects.