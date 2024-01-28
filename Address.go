package blockchain

import (
	"context"
	gohttpclient "github.com/bozd4g/go-http-client"
)

type AddressBalance struct {
	Balance                   int64   `json:"balance"`
	ReceivedAmount            int64   `json:"receivedAmount"`
	ReceivedTxCount           int     `json:"receivedTxCount"`
	SentAmount                int64   `json:"sentAmount"`
	SentTxCount               int     `json:"sentTxCount"`
	FirstReceivedTxPointer    string  `json:"firstReceivedTxPointer"`
	FirstSentTxPointer        string  `json:"firstSentTxPointer"`
	LastTxPointer             string  `json:"lastTxPointer"`
	LargestSpentTxAmount      int64   `json:"largestSpentTxAmount"`
	LargestSpentTxPointer     string  `json:"largestSpentTxPointer"`
	LargestReceivedTxAmount   int64   `json:"largestReceivedTxAmount"`
	LargestReceivedTxPointer  string  `json:"largestReceivedTxPointer"`
	ReceivedOutsCount         int     `json:"receivedOutsCount"`
	SpentOutsCount            int     `json:"spentOutsCount"`
	PendingReceivedAmount     int64   `json:"pendingReceivedAmount"`
	PendingSentAmount         int64   `json:"pendingSentAmount"`
	PendingReceivedTxCount    int     `json:"pendingReceivedTxCount"`
	PendingSentTxCount        int     `json:"pendingSentTxCount"`
	PendingReceivedOutsCount  int     `json:"pendingReceivedOutsCount"`
	PendingSpentOutsCount     int     `json:"pendingSpentOutsCount"`
	Type                      string  `json:"type"`
}


type AddressTransactions struct {
	List  []AddressTransaction `json:"list"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
	Pages int                  `json:"pages"`
	Time  int64                `json:"time"`
}

type AddressTransaction struct {
	TxID   string `json:"txId"`
	VOut   int    `json:"vOut"`
	Block  int    `json:"block"`
	TxIndex int   `json:"txIndex"`
	Amount int64  `json:"amount"`
}

type AddressUTXO struct {
	TxID   string `json:"txId"`
	VOut   int    `json:"vOut"`
	Block  int    `json:"block"`
	TxIndex int   `json:"txIndex"`
	Amount int64  `json:"amount"`
}

type AddressUTXO struct {
	TxID    string `json:"txId"`
	VOut    int    `json:"vOut"`
	Block   int    `json:"block"`
	TxIndex int    `json:"txIndex"`
	Amount  int64  `json:"amount"`
}

type AddressUTXOs []AddressUTXO

type Address interface {
  ApiUrl string
  GetAddressBalance(string address) (AddressBalance, error)
	GetAddressTransactions(string address) (AddressTransaction, error)
	GetAddressUTXO(string address) (AddressUTXOs, error)
	getUnconfirmedAddressTransactions(string address) (AddressTransactions, error)
}

func (a *Address) GetAddressBalance(string address) (AddressBalance, error)  {
	ctx := context.Background()
	client := gohttpclient.New(API_URL)
	response, err := client.Get(ctx, address + "/")
	if err != nil {
	  return nill, err	
  }

	var AddressBalance addressBalance

	if err := response.Unmarshall(&addressBalance); err != nil {
		return nil, err
	}

	return addressBalance, nil
}