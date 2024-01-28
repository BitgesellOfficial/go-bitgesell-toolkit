package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NewTransaction creates a new instance of the Transaction SDK
func NewTransaction(config SDKConfig) *Transaction {
	apiURL := config.BaseAPIURL
	if apiURL == "" {
		apiURL = "https://api.bitaps.com/bgl/v1/blockchain"
	}

	return &Transaction{
		apiV1URL: apiURL,
		logger:   func(arg string) { fmt.Println(arg) }, // Update with your logger implementation
	}
}

// GetTransactionByHash fetches a transaction by its hash
func (t *Transaction) GetTransactionByHash(txHash string) (*AddressTransaction, error) {
	url := fmt.Sprintf("%s/transaction/%s", t.apiV1URL, txHash)
	data, err := t.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into AddressTransaction struct
	var addressTransaction AddressTransaction
	if err := json.Unmarshal(data, &addressTransaction); err != nil {
		return nil, err
	}

	return &addressTransaction, nil
}

// GetTransactionMerkleProof fetches the MerkelProof of a transaction
func (t *Transaction) GetTransactionMerkleProof(txHash string) (*TransactionMerkleProof, error) {
	url := fmt.Sprintf("%s/transaction/merkle_proof/%s", t.apiV1URL, txHash)
	data, err := t.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into TransactionMerkleProof struct
	var transactionMerkleProof TransactionMerkleProof
	if err := json.Unmarshal(data, &transactionMerkleProof); err != nil {
		return nil, err
	}

	return &transactionMerkleProof, nil
}

// Other methods...

// get is a private method to make GET requests
func (t *Transaction) get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		t.logger(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.logger(err.Error())
		return nil, err
	}

	return body, nil
}
