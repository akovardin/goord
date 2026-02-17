//nolint:errcheck
package ord

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCIDList(t *testing.T) {
	testResponse := CIDListResponse{
		CIDs:            []string{"cid1", "cid2", "cid3"},
		TotalItemsCount: 3,
		Limit:           10,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")
		assert.Equal(t, "/v1/cid", r.URL.Path, "Expected path /v1/cid")
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

	result, err := client.GetCIDList(context.Background(), 0, 10)
	require.NoError(t, err, "GetCIDList should not return an error")

	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.CIDs, result.CIDs, "CIDs should match")
}

func TestClient_GetCIDList_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetCIDList(context.Background(), 0, 10)
	require.Error(t, err, "GetCIDList should return an error")

	assert.Contains(t, err.Error(), "failed to get CID list", "Error message should contain expected text")
}

func TestClient_GetCID(t *testing.T) {
	testCID := CID{
		CID:  "test-cid",
		Name: "Test CID",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")
		assert.Equal(t, "/v1/cid/test-cid", r.URL.Path, "Expected path /v1/cid/test-cid")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testCID)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetCID(context.Background(), "test-cid")
	require.NoError(t, err, "GetCID should not return an error")

	assert.Equal(t, testCID.CID, result.CID, "CID should match")
	assert.Equal(t, testCID.Name, result.Name, "Name should match")
}

func TestClient_GetCID_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("CID not found"))
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetCID(context.Background(), "non-existent-cid")
	require.Error(t, err, "GetCID should return an error")

	assert.Contains(t, err.Error(), "failed to get CID", "Error message should contain expected text")
}

func TestClient_CreateCID(t *testing.T) {
	testCID := CID{
		CID:  "test-cid",
		Name: "Test CID",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected PUT request")
		assert.Equal(t, "/v1/cid/test-cid", r.URL.Path, "Expected path /v1/cid/test-cid")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	err := client.CreateCID(context.Background(), "test-cid", testCID)
	require.NoError(t, err, "CreateCID should not return an error")
}

func TestClient_CreateCID_Error(t *testing.T) {
	testCID := CID{
		CID:  "test-cid",
		Name: "Test CID",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	err := client.CreateCID(context.Background(), "test-cid", testCID)
	require.Error(t, err, "CreateCID should return an error")

	assert.Contains(t, err.Error(), "failed to create CID", "Error message should contain expected text")
}
