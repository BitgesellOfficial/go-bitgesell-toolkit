package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAddressBalance(t *testing.T) {
	mockServer := createMockServer("address/state/mockAddress", `{"mock": "data"}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL,
	}
	addressAPI := NewAddress(config)

	_, err := addressAPI.GetAddressBalance("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetAddressTransactions(t *testing.T) {
	mockServer := createMockServer("address/transactions/mockAddress", `{"mock": "data"}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL,
	}
	addressAPI := NewAddress(config)

	_, err := addressAPI.GetAddressTransactions("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetUnconfirmedAddressTransactions(t *testing.T) {
	mockServer := createMockServer("address/unconfirmed/transactions/mockAddress", `{"mock": "data"}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL,
	}
	addressAPI := NewAddress(config)

	_, err := addressAPI.GetUnconfirmedAddressTransactions("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetAddressUTXO(t *testing.T) {
	mockServer := createMockServer("address/unconfirmed/transactions/mockAddress", `{"mock": "data"}`)
	defer mockServer.Close()

	config := SDKConfig{
		BaseAPIURL: mockServer.URL,
	}
	addressAPI := NewAddress(config)

	_, err := addressAPI.GetAddressUTXO("mockAddress")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func createMockServer(path, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bgl/v1/blockchain/"+path {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}))
}
