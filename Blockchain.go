package blockchain

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

// Blockchain represents the blockchain SDK
type Blockchain struct {
	apiV1URL string
	logger   func(string)
}

// NewBlockchain creates a new instance of the Blockchain SDK
func NewBlockchain(config SDKConfig) *Blockchain {
	apiURL := config.BaseAPIURL
	if apiURL == "" {
		apiURL = "https://api.bitaps.com/bgl/v1/blockchain"
	}

	return &Blockchain{
		apiV1URL: apiURL,
		logger:   func(arg string) { fmt.Println(arg) }, // Update with your logger implementation
	}
}

// GetBlockByHash fetches a given block by a given block hash
func (b *Blockchain) GetBlockByHash(blockHash string) (*Block, error) {
	url := fmt.Sprintf("%s/block/%s", b.apiV1URL, blockHash)
	data, err := b.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into Block struct
	var block Block
	if err := json.Unmarshal(data, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetBlockByHeight fetches a given block by a given block height
func (b *Blockchain) GetBlockByHeight(blockHeight int) (*Block, error) {
	url := fmt.Sprintf("%s/block/%d", b.apiV1URL, blockHeight)
	data, err := b.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into Block struct
	var block Block
	if err := json.Unmarshal(data, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetLastBlocksByCount fetches the last N blocks
func (b *Blockchain) GetLastBlocksByCount(blockCount int) ([]Block, error) {
	url := fmt.Sprintf("%s/blocks/last/%d", b.apiV1URL, blockCount)
	data, err := b.get(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal data into []Block
	var blocks []Block
	if err := json.Unmarshal(data, &blocks); err != nil {
		return nil, err
	}

	return blocks, nil
}

// Other methods...

// get is a private method to make GET requests
func (b *Blockchain) get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		b.logger(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		b.logger(err.Error())
		return nil, err
	}

	return body, nil
}
