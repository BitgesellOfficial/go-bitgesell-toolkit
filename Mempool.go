package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SDKConfig represents the configuration for the SDK
type SDKConfig struct {
	BaseAPIURL string // Add any other fields as needed
}

// Mempool represents the mempool SDK
type Mempool struct {
	apiV1URL string
	logger   func(string)
}

// NewMempool creates a new instance of the Mempool SDK
func NewMempool(config SDKConfig) *Mempool {
	apiURL := config.BaseAPIURL
	if apiURL == "" {
		apiURL = "https://api.bitaps.com/bgl/v1/blockchain"
	}

	return &Mempool{
		apiV1URL: apiURL,
		logger:   func(arg string) { fmt.Println(arg) }, // Update with your logger implementation
	}
}

// GetMempoolTransactions fetches transactions in the mempool
func (m *Mempool) GetMempoolTransactions(limit int, order string, fromTimestamp int, page int) (*Mempool, error) {
	url := fmt.Sprintf("%s/mempool/transactions?page=%d&limit=%d&order=%s&from_timestamp=%d", m.apiV1URL, page, limit, order, fromTimestamp)
	data, err := m.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into Mempool struct
	var mempool Mempool
	if err := json.Unmarshal(data, &mempool); err != nil {
		return nil, err
	}

	return &mempool, nil
}

// GetMempoolState fetches the state of the mempool
func (m *Mempool) GetMempoolState(limit int, order string, fromTimestamp int, page int) (*MempoolState, error) {
	url := fmt.Sprintf("%s/mempool/state?page=%d&limit=%d&order=%s&from_timestamp=%d", m.apiV1URL, page, limit, order, fromTimestamp)
	data, err := m.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into MempoolState struct
	var mempoolState MempoolState
	if err := json.Unmarshal(data, &mempoolState); err != nil {
		return nil, err
	}

	return &mempoolState, nil
}

// Other methods...

// get is a private method to make GET requests
func (m *Mempool) get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		m.logger(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.logger(err.Error())
		return nil, err
	}

	return body, nil
}
