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

func TestClient_CreativeMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v3/creative" && r.Method == "GET":
			handleGetCreatives(w, r)
		case r.URL.Path == "/v3/creative/list/erids" && r.Method == "GET":
			handleGetCreativeERIDs(w, r)
		case r.URL.Path == "/v3/creative/list/erid_external_ids" && r.Method == "GET":
			handleGetCreativeERIDExternalIDPairs(w, r)
		case r.URL.Path == "/v2/creative/test-external-id" && r.Method == "PUT":
			handleCreateCreativeV2(w, r)
		case r.URL.Path == "/v2/creative/test-external-id" && r.Method == "GET":
			handleGetCreativeV2(w, r)
		case r.URL.Path == "/v2/creative/by_erid/test-erid" && r.Method == "GET":
			handleGetCreativeByERIDV2(w, r)
		case r.URL.Path == "/v3/creative/test-external-id" && r.Method == "PUT":
			handleCreateCreativeV3(w, r)
		case r.URL.Path == "/v3/creative/test-external-id" && r.Method == "GET":
			handleGetCreativeV3(w, r)
		case r.URL.Path == "/v3/creative/by_erid/test-erid" && r.Method == "GET":
			handleGetCreativeByERIDV3(w, r)
		case r.URL.Path == "/v3/creative/test-external-id/add_text" && r.Method == "POST":
			handleAddTextsToCreative(w, r)
		case r.URL.Path == "/v3/creative/test-external-id/add_media" && r.Method == "POST":
			handleAddMediaToCreative(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error": "Not found"}`)
		}
	}))
	defer server.Close()

	client, _ := NewClient(
		WithBase(server.URL),
		WithToken("test-token"),
	)

	t.Run("GetCreatives", func(t *testing.T) {
		response, err := client.GetCreatives(context.Background(), 0, 10)
		require.NoError(t, err)
		assert.Equal(t, []string{"ext1", "ext2"}, response.ExternalIDs)
		assert.Equal(t, 2, response.TotalItemsCount)
		assert.Equal(t, 10, response.Limit)
	})

	t.Run("GetCreativeERIDs", func(t *testing.T) {
		response, err := client.GetCreativeERIDs(context.Background(), 0, 10)
		require.NoError(t, err)
		assert.Equal(t, []string{"erid1", "erid2"}, response.ERIDs)
		assert.Equal(t, 2, response.TotalItemsCount)
		assert.Equal(t, 10, response.Limit)
	})

	t.Run("GetCreativeERIDExternalIDPairs", func(t *testing.T) {
		response, err := client.GetCreativeERIDExternalIDPairs(context.Background(), 0, 10)
		require.NoError(t, err)
		expectedItems := []CreativeERIDExternalIDPair{
			{ERID: "erid1", ExternalID: "ext1"},
			{ERID: "erid2", ExternalID: "ext2"},
		}
		assert.Equal(t, expectedItems, response.Items)
		assert.Equal(t, 2, response.TotalItemsCount)
		assert.Equal(t, 10, response.Limit)
	})

	t.Run("CreateCreativeV2", func(t *testing.T) {
		creative := CreateCreativeV2Request{
			Name:  stringPtr("Test Creative"),
			Form:  "video",
			KKTUs: []string{"12345"},
		}
		err := client.CreateCreativeV2(context.Background(), "test-external-id", creative)
		require.NoError(t, err)
	})

	t.Run("GetCreativeV2", func(t *testing.T) {
		creative, err := client.GetCreativeV2(context.Background(), "test-external-id")
		require.NoError(t, err)
		assert.Equal(t, "test-erid", creative.ERID)
		assert.Equal(t, "Test Creative", *creative.Name)
	})

	t.Run("GetCreativeByERIDV2", func(t *testing.T) {
		creative, err := client.GetCreativeByERIDV2(context.Background(), "test-erid")
		require.NoError(t, err)
		assert.Equal(t, "test-erid", creative.ERID)
		assert.Equal(t, "Test Creative", *creative.Name)
	})

	t.Run("CreateCreativeV3", func(t *testing.T) {
		creative := CreateCreativeV3Request{
			Name:  stringPtr("Test Creative V3"),
			Form:  "video",
			KKTUs: []string{"12345"},
		}
		err := client.CreateCreativeV3(context.Background(), "test-external-id", creative)
		require.NoError(t, err)
	})

	t.Run("GetCreativeV3", func(t *testing.T) {
		creative, err := client.GetCreativeV3(context.Background(), "test-external-id")
		require.NoError(t, err)
		assert.Equal(t, "test-erid", creative.ERID)
		assert.Equal(t, "Test Creative V3", *creative.Name)
	})

	t.Run("GetCreativeByERIDV3", func(t *testing.T) {
		creative, err := client.GetCreativeByERIDV3(context.Background(), "test-erid")
		require.NoError(t, err)
		assert.Equal(t, "test-erid", creative.ERID)
		assert.Equal(t, "Test Creative V3", *creative.Name)
	})

	t.Run("AddTextsToCreative", func(t *testing.T) {
		texts := []string{"Text 1", "Text 2"}
		err := client.AddTextsToCreative(context.Background(), "test-external-id", texts)
		require.NoError(t, err)
	})

	t.Run("AddMediaToCreative", func(t *testing.T) {
		mediaIDs := []string{"media1", "media2"}
		err := client.AddMediaToCreative(context.Background(), "test-external-id", mediaIDs)
		require.NoError(t, err)
	})
}

func stringPtr(s string) *string {
	return &s
}

func handleGetCreatives(w http.ResponseWriter, r *http.Request) {
	response := CreativeListResponse{
		ExternalIDs:     []string{"ext1", "ext2"},
		TotalItemsCount: 2,
		Limit:           10,
	}
	json.NewEncoder(w).Encode(response)
}

func handleGetCreativeERIDs(w http.ResponseWriter, r *http.Request) {
	response := CreativeERIDsListResponse{
		ERIDs:           []string{"erid1", "erid2"},
		TotalItemsCount: 2,
		Limit:           10,
	}
	json.NewEncoder(w).Encode(response)
}

func handleGetCreativeERIDExternalIDPairs(w http.ResponseWriter, r *http.Request) {
	response := CreativeERIDExternalIDPairsResponse{
		Items: []CreativeERIDExternalIDPair{
			{ERID: "erid1", ExternalID: "ext1"},
			{ERID: "erid2", ExternalID: "ext2"},
		},
		TotalItemsCount: 2,
		Limit:           10,
	}
	json.NewEncoder(w).Encode(response)
}

func handleCreateCreativeV2(w http.ResponseWriter, r *http.Request) {
	var creative CreateCreativeV2Request
	if err := json.NewDecoder(r.Body).Decode(&creative); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGetCreativeV2(w http.ResponseWriter, r *http.Request) {
	creative := Creative{
		ERID:  "test-erid",
		Name:  stringPtr("Test Creative"),
		Form:  "video",
		KKTUs: []string{"12345"},
	}
	json.NewEncoder(w).Encode(creative)
}

func handleGetCreativeByERIDV2(w http.ResponseWriter, r *http.Request) {
	creative := Creative{
		ERID:  "test-erid",
		Name:  stringPtr("Test Creative"),
		Form:  "video",
		KKTUs: []string{"12345"},
	}
	json.NewEncoder(w).Encode(creative)
}

func handleCreateCreativeV3(w http.ResponseWriter, r *http.Request) {
	var creative CreateCreativeV3Request
	if err := json.NewDecoder(r.Body).Decode(&creative); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGetCreativeV3(w http.ResponseWriter, r *http.Request) {
	creative := Creative{
		ERID:  "test-erid",
		Name:  stringPtr("Test Creative V3"),
		Form:  "video",
		KKTUs: []string{"12345"},
	}
	json.NewEncoder(w).Encode(creative)
}

func handleGetCreativeByERIDV3(w http.ResponseWriter, r *http.Request) {
	creative := Creative{
		ERID:  "test-erid",
		Name:  stringPtr("Test Creative V3"),
		Form:  "video",
		KKTUs: []string{"12345"},
	}
	json.NewEncoder(w).Encode(creative)
}

func handleAddTextsToCreative(w http.ResponseWriter, r *http.Request) {
	var req AddTextsToCreativeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleAddMediaToCreative(w http.ResponseWriter, r *http.Request) {
	var req AddMediaToCreativeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
}
