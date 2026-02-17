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

func TestClient_GetPads(t *testing.T) {
	testResponse := PadListResponse{
		ExternalIDs:     []string{"pad1", "pad2"},
		TotalItemsCount: 2,
		Limit:           100,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		assert.Equal(t, "/v1/pad", r.URL.Path, "Expected path /v1/pad")

		offset := r.URL.Query().Get("offset")
		limit := r.URL.Query().Get("limit")
		personExternalID := r.URL.Query().Get("person_external_id")
		assert.Equal(t, "0", offset, "Expected offset=0")
		assert.Equal(t, "100", limit, "Expected limit=100")
		assert.Equal(t, "", personExternalID, "Expected person_external_id to be empty")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetPads(context.Background(), 0, 100, "")
	require.NoError(t, err, "GetPads should not return an error")

	assert.Equal(t, testResponse.TotalItemsCount, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, testResponse.Limit, result.Limit, "Limit should match")
	assert.Equal(t, testResponse.ExternalIDs, result.ExternalIDs, "ExternalIDs should match")
}

func TestClient_GetPads_WithPersonExternalID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		personExternalID := r.URL.Query().Get("person_external_id")
		assert.Equal(t, "person1", personExternalID, "Expected person_external_id=person1")

		response := PadListResponse{
			ExternalIDs:     []string{"pad1", "pad2"},
			TotalItemsCount: 2,
			Limit:           100,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetPads(context.Background(), 0, 100, "person1")
	require.NoError(t, err, "GetPads should not return an error")

	assert.Equal(t, 2, result.TotalItemsCount, "TotalItemsCount should match")
	assert.Equal(t, 100, result.Limit, "Limit should match")
	assert.Equal(t, []string{"pad1", "pad2"}, result.ExternalIDs, "ExternalIDs should match")
}

func TestClient_GetPads_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetPads(context.Background(), 0, 100, "")
	require.Error(t, err, "GetPads should return an error")

	assert.Contains(t, err.Error(), "failed to get pads", "Error message should contain expected text")
}

func TestClient_GetRestrictedPads(t *testing.T) {
	testURLs := []string{"restricted.com", "limited.com"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		assert.Equal(t, "/v1/pad/info/restricted", r.URL.Path, "Expected path /v1/pad/info/restricted")

		response := struct {
			URLs []string `json:"urls"`
		}{
			URLs: testURLs,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	urls, err := client.GetRestrictedPads(context.Background())
	require.NoError(t, err, "GetRestrictedPads should not return an error")

	assert.Equal(t, testURLs, urls, "URLs should match")
}

func TestClient_GetRestrictedPads_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetRestrictedPads(context.Background())
	require.Error(t, err, "GetRestrictedPads should return an error")

	assert.Contains(t, err.Error(), "failed to get restricted pads", "Error message should contain expected text")
}

func TestClient_GetPad(t *testing.T) {
	testPad := Pad{
		PersonExternalID: "person1",
		IsOwner:          true,
		Type:             "web",
		Name:             "Test Pad",
		URL:              padStringPtr("https://test.com"),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected GET request")

		assert.Equal(t, "/v1/pad/test-pad", r.URL.Path, "Expected path /v1/pad/test-pad")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testPad)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	result, err := client.GetPad(context.Background(), "test-pad")
	require.NoError(t, err, "GetPad should not return an error")

	assert.Equal(t, testPad.PersonExternalID, result.PersonExternalID, "PersonExternalID should match")
	assert.Equal(t, testPad.IsOwner, result.IsOwner, "IsOwner should match")
	assert.Equal(t, testPad.Type, result.Type, "Type should match")
	assert.Equal(t, testPad.Name, result.Name, "Name should match")
	assert.Equal(t, *testPad.URL, *result.URL, "URL should match")
}

func TestClient_GetPad_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	_, err := client.GetPad(context.Background(), "test-pad")
	require.Error(t, err, "GetPad should return an error")

	assert.Contains(t, err.Error(), "failed to get pad", "Error message should contain expected text")
}

func TestClient_CreatePad(t *testing.T) {
	testPad := Pad{
		PersonExternalID: "person1",
		IsOwner:          true,
		Type:             "web",
		Name:             "New Pad",
		URL:              padStringPtr("https://new.com"),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected PUT request")

		assert.Equal(t, "/v1/pad/test-pad", r.URL.Path, "Expected path /v1/pad/test-pad")

		var pad Pad
		err := json.NewDecoder(r.Body).Decode(&pad)
		require.NoError(t, err, "Should be able to decode request body")
		assert.Equal(t, testPad.PersonExternalID, pad.PersonExternalID, "PersonExternalID should match")
		assert.Equal(t, testPad.IsOwner, pad.IsOwner, "IsOwner should match")
		assert.Equal(t, testPad.Type, pad.Type, "Type should match")
		assert.Equal(t, testPad.Name, pad.Name, "Name should match")
		assert.Equal(t, *testPad.URL, *pad.URL, "URL should match")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	err := client.CreatePad(context.Background(), "test-pad", testPad)
	require.NoError(t, err, "CreatePad should not return an error")
}

func TestClient_CreatePad_Error(t *testing.T) {
	testPad := Pad{
		PersonExternalID: "person1",
		IsOwner:          true,
		Type:             "web",
		Name:             "New Pad",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	err := client.CreatePad(context.Background(), "test-pad", testPad)
	require.Error(t, err, "CreatePad should return an error")

	assert.Contains(t, err.Error(), "failed to create pad", "Error message should contain expected text")
}

func padStringPtr(s string) *string {
	return &s
}
