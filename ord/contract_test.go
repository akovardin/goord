package ord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetContracts(t *testing.T) {
	testResponse := ContractListResponse{
		ExternalIDs:     []string{"contract1", "contract2", "contract3"},
		TotalItemsCount: 3,
		Limit:           10,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")
		assert.Equal(t, "/v1/contract", r.URL.Path, "Expected path /v1/contract")
		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")
		assert.Equal(t, "0", offset, "Expected offset=0")
		assert.Equal(t, "10", limit, "Expected limit=10")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetContracts(context.Background(), 0, 10)
	require.NoError(t, err, "GetContracts should not return an error")

	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.ExternalIDs, result.ExternalIDs, "ExternalIDs should match")
}

func TestClient_GetContracts_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetContracts(context.Background(), 0, 10)
	require.Error(t, err, "GetContracts should return an error")

	assert.Contains(t, err.Error(), "failed to get contracts", "Error message should contain expected text")
}

func TestClient_GetContract(t *testing.T) {
	testContract := Contract{
		CreateDate:             "2023-07-04T11:27:40Z",
		Type:                   "service",
		ClientExternalID:       "client1",
		ContractorExternalID:   "contractor1",
		SubjectType:            "org_distribution",
		Date:                   "2022-12-01",
		HasAdditionalContracts: false,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")
		assert.Equal(t, "/v1/contract/contract1", r.URL.Path, "Expected path /v1/contract/contract1")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testContract)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetContract(context.Background(), "contract1")
	require.NoError(t, err, "GetContract should not return an error")

	assert.Equal(t, testContract.CreateDate, result.CreateDate, "CreateDate should match")
	assert.Equal(t, testContract.Type, result.Type, "Type should match")
	assert.Equal(t, testContract.ClientExternalID, result.ClientExternalID, "ClientExternalID should match")
	assert.Equal(t, testContract.ContractorExternalID, result.ContractorExternalID, "ContractorExternalID should match")
	assert.Equal(t, testContract.SubjectType, result.SubjectType, "SubjectType should match")
	assert.Equal(t, testContract.Date, result.Date, "Date should match")
}

func TestClient_CreateContract(t *testing.T) {
	testContract := CreateContractRequest{
		Type:                 "service",
		ClientExternalID:     "client1",
		ContractorExternalID: "contractor1",
		SubjectType:          "org_distribution",
		Date:                 "2022-12-01",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected PUT request")
		assert.Equal(t, "/v1/contract/contract1", r.URL.Path, "Expected path /v1/contract/contract1")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	err := client.CreateContract(context.Background(), "contract1", testContract)
	require.NoError(t, err, "CreateContract should not return an error")
}

func TestClient_RequestCID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "Expected POST request")
		assert.Equal(t, "/v1/contract/contract1/create_cid", r.URL.Path, "Expected path /v1/contract/contract1/create_cid")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	err := client.RequestCID(context.Background(), "contract1")
	require.NoError(t, err, "RequestCID should not return an error")
}
