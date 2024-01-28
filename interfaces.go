package blockchain

// Block represents the block type
type Block struct {
	Height             int    `json:"height"`
	Hash               string `json:"hash"`
	Header             string `json:"header"`
	AdjustedTimestamp  int    `json:"adjustedTimestamp"`
}

// SDKConfig represents the configuration for the SDK
type SDKConfig struct {
	BaseAPIURL string `json:"baseAPIURL"`
	Logger     func(string) `json:"-"`
	APIKey     string `json:"apiKey,omitempty"`
}

// BlockHeader represents the block header type
type BlockHeader struct {
	Data []interface{} `json:"data"`
	Time int64          `json:"time"`
}

// BlockStats represents the block statistics type
type BlockStats struct {
	Height              int     `json:"height"`
	Hash                string  `json:"hash"`
	Header              string  `json:"header"`
	Version             int     `json:"version"`
	PreviousBlockHash   string  `json:"previousBlockHash"`
	MerkleRoot          string  `json:"merkleRoot"`
	Bits                int     `json:"bits"`
	Nonce               int     `json:"nonce"`
	Weight              int     `json:"weight"`
	Size                int     `json:"size"`
	StrippedSize        int     `json:"strippedSize"`
	Amount              int     `json:"amount"`
	Target              string  `json:"target"`
	Miner               string  `json:"miner"`
	MedianBlockTime     int     `json:"medianBlockTime"`
	BlockTime           int     `json:"blockTime"`
	ReceivedTimestamp   int     `json:"receivedTimestamp"`
	AdjustedTimestamp   int     `json:"adjustedTimestamp"`
	BitsHex             string  `json:"bitsHex"`
	NonceHex            string  `json:"nonceHex"`
	VersionHex          string  `json:"versionHex"`
	Difficulty          int     `json:"difficulty"`
	BlockDifficulty     int     `json:"blockDifficulty"`
	NextBlockHash       string  `json:"nextBlockHash"`
	EstimatedBlockReward int     `json:"estimatedBlockReward"`
	BlockReward         int     `json:"blockReward"`
	BlockFeeReward      int     `json:"blockFeeReward"`
	Confirmations       int     `json:"confirmations"`
	TransactionsCount   int     `json:"transactionsCount"`
	Coinbase            string  `json:"coinbase"`
	Statistics          object  `json:"statistics"`
	Time                int64   `json:"time"`
}

// Transaction represents the transaction type
type Transaction struct {
	Segwit             bool    `json:"segwit"`
	RBF                bool    `json:"rbf"`
	TxID               string  `json:"txId"`
	Hash               string  `json:"hash"`
	Version            int     `json:"version"`
	Size               int     `json:"size"`
	VSize              int     `json:"vSize"`
	BSize              int     `json:"bSize"`
	LockTime           int     `json:"lockTime"`
	VIn                object  `json:"vIn"`
	VOut               object  `json:"vOut"`
	Confirmations      int     `json:"confirmations"`
	BlockIndex         int     `json:"blockIndex"`
	Coinbase           bool    `json:"coinbase"`
	Data               string  `json:"data"`
	RawTx              string  `json:"rawTx"`
	Amount             int     `json:"amount"`
	Flag               string  `json:"flag"`
	Weight             int     `json:"weight"`
	Timestamp          int64   `json:"timestamp"`
	MerkleProof        string  `json:"merkleProof"`
	InputsAmount       int     `json:"inputsAmount"`
	OutputAddresses    int     `json:"outputAddresses"`
	InputAddresses     int     `json:"inputAddresses"`
	Fee                int     `json:"fee"`
	OutputsAmount      int     `json:"outputsAmount"`
	Inputs             int     `json:"inputs"`
	Outputs            int     `json:"outputs"`
}

// Transactions represents the transactions type
type Transactions struct {
	List   []Transaction `json:"list"`
	Page   int           `json:"page"`
	Pages  int           `json:"pages"`
	Total  int           `json:"total"`
	Limit  int           `json:"limit"`
	Time   int64         `json:"time"`
}

// UTXO represents the Unspent Transaction Output type
type UTXO struct {
	TxID     string `json:"txId"`
	VOut     int    `json:"vOut"`
	TxIndex  int    `json:"txIndex"`
	Amount   int    `json:"amount"`
	Address  string `json:"address"`
	Script   string `json:"script"`
	Type     string `json:"type"`
	Time     int64  `json:"time"`
}

// UTXOs represents an array of UTXO
type UTXOs []UTXO

// MempoolTxes represents the mempool transactions type
type MempoolTxes struct {
	List           []interface{} `json:"list"`
	Page           int           `json:"page"`
	Limit          int           `json:"limit"`
	Pages          int           `json:"pages"`
	Count          int           `json:"count"`
	FromTimestamp  int64         `json:"fromTimestamp"`
	Time           int64         `json:"time"`
}

// MempoolState represents the mempool state type
type MempoolState struct {
	Inputs       object `json:"inputs"`
	Outputs      object `json:"outputs"`
	Transactions object `json:"transactions"`
	Time         int64  `json:"time"`
}

// MempoolRecommendedFee represents the recommended fee from the mempool type
type MempoolRecommendedFee struct {
	Best      int   `json:"best"`
	Best4h    int   `json:"best4h"`
	BestHourly int   `json:"bestHourly"`
	Time      int64 `json:"time"`
}

// AddressBalance represents the address balance type
type AddressBalance struct {
	Balance                  int    `json:"balance"`
	ReceivedAmount           int    `json:"receivedAmount"`
	ReceivedTxCount          int    `json:"receivedTxCount"`
	SentAmount               int    `json:"sentAmount"`
	SentTxCount              int    `json:"sentTxCount"`
	FirstReceivedTxPointer   string `json:"firstReceivedTxPointer"`
	FirstSentTxPointer       string `json:"firstSentTxPointer"`
	LastTxPointer            string `json:"lastTxPointer"`
	LargestSpentTxAmount     int    `json:"largestSpentTxAmount"`
	LargestSpentTxPointer    string `json:"largestSpentTxPointer"`
	LargestReceivedTxAmount  int    `json:"largestReceivedTxAmount"`
	LargestReceivedTxPointer string `json:"largestReceivedTxPointer"`
	ReceivedOutsCount        int    `json:"receivedOutsCount"`
	SpentOutsCount           int    `json:"spentOutsCount"`
	PendingReceivedAmount    int    `json:"pendingReceivedAmount"`
	PendingSentAmount        int    `json:"pendingSentAmount"`
	PendingReceivedTxCount   int    `json:"pendingReceivedTxCount"`
	PendingSentTxCount       int    `json:"pendingSentTxCount"`
	PendingReceivedOutsCount int    `json:"pendingReceivedOutsCount"`
	PendingSpentOutsCount    int    `json:"pendingSpentOutsCount"`
	Type                     string `json:"type"`
}

// AddressTransaction represents the address transaction type
type AddressTransaction struct {
	Segwit              bool    `json:"segwit"`
	RBF                 bool    `json:"rbf"`
	TxID                string  `json:"txId"`
	Version             int     `json:"version"`
	Size                int     `json:"size"`
	VSize               int     `json:"vSize"`
	BSize               int     `json:"bSize"`
	LockTime            int     `json:"lockTime"`
	VIn                 object  `json:"vIn"`
	VOut                object  `json:"vOut"`
	Confirmations       int     `json:"confirmations"`
	BlockTime           int     `json:"blockTime"`
	BlockIndex          int     `json:"blockIndex"`
	Coinbase            bool    `json:"coinbase"`
	Fee                 int     `json:"fee"`
	Data                string  `json:"data"`
	Amount              int     `json:"amount"`
	Weight              int     `json:"weight"`
	BlockHeight         int     `json:"blockHeight"`
	Timestamp           int64   `json:"timestamp"`
	InputsAmount        int     `json:"inputsAmount"`
	InputAddressCount   int     `json:"inputAddressCount"`
	OutAddressCount     int     `json:"outAddressCount"`
	InputsCount         int     `json:"inputsCount"`
	OutsCount           int     `json:"outsCount"`
	OutputsAmount       int     `json:"outputsAmount"`
	AddressReceived     int     `json:"addressReceived"`
	AddressOuts         int     `json:"addressOuts"`
	AddressSent         int     `json:"addressSent"`
	AddressInputs       int     `json:"addressInputs"`
}

// AddressTransactions represents the address transactions type
type AddressTransactions struct {
	List  []AddressTransaction `json:"list"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
	Pages int                  `json:"pages"`
	Time  int64                `json:"time"`
}

// AddressUTXO represents the address UTXO type
type AddressUTXO struct {
	TxID    string `json:"txId"`
	VOut    int    `json:"vOut"`
	Block   int    `json:"block"`
	TxIndex int    `json:"txIndex"`
	Amount  int    `json:"amount"`
}

// AddressUTXOs represents an array of address UTXO
type AddressUTXOs []AddressUTXO

// TransactionMerkelProof represents the transaction MerkelProof type
type TransactionMerkelProof struct {
	BlockHeight int    `json:"blockHeight"`
	BlockIndex  int    `json:"blockIndex"`
	MerkleProof string `json:"merkleProof"`
	Time        int64  `json:"time"`
}
