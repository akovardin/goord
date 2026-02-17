//nolint:errcheck
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

func TestClient_GetPersons(t *testing.T) {
	testResponse := PersonListResponse{
		ExternalIDs:     []string{"id1", "id2", "id3"},
		TotalItemsCount: 3,
		Limit:           10,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/person", r.URL.Path)
		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")
		assert.Equal(t, "0", offset)
		assert.Equal(t, "10", limit)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetPersons(context.Background(), 0, 10)
	require.NoError(t, err)

	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount)
	assert.Equal(t, testResponse.Limit, result.Limit)
	assert.Equal(t, testResponse.ExternalIDs, result.ExternalIDs)
}

func TestClient_GetPersons_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetPersons(context.Background(), 0, 10)
	require.Error(t, err)

	assert.Contains(t, err.Error(), "failed to get persons")
}
