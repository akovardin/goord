package ord

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_MediaMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v1/media" && r.Method == "GET":
			handleGetMediaList(w, r)
		case strings.HasPrefix(r.URL.Path, "/v1/media/") && strings.HasSuffix(r.URL.Path, "/info") && r.Method == "GET":
			handleGetMediaInfo(w, r)
		case r.URL.Path == "/v1/get_media_info" && r.Method == "POST":
			handleGetMediaInfoBatch(w, r)
		case strings.HasPrefix(r.URL.Path, "/v1/media/") && !strings.HasSuffix(r.URL.Path, "/info") && r.Method == "GET":
			handleGetMediaBinary(w, r)
		case strings.HasPrefix(r.URL.Path, "/v1/media/") && r.Method == "PUT":
			handleUploadMedia(w, r)
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

	t.Run("GetMediaList", func(t *testing.T) {
		response, err := client.GetMediaList(context.Background(), 0, 100)
		require.NoError(t, err)
		assert.Equal(t, []string{"media1", "media2"}, response.ExternalIDs)
		assert.Equal(t, 2, response.TotalItemsCount)
		assert.Equal(t, 100, response.Limit)
	})

	t.Run("UploadMedia", func(t *testing.T) {
		fileContent := "test file content"
		reader := strings.NewReader(fileContent)
		sha256, err := client.UploadMedia(context.Background(), "test-media", "test.txt", reader)
		require.NoError(t, err)
		assert.Equal(t, "test-sha256", *sha256)
	})

	t.Run("GetMediaBinary", func(t *testing.T) {
		data, err := client.GetMediaBinary(context.Background(), "test-media")
		require.NoError(t, err)
		assert.Equal(t, []byte("test binary data"), data)
	})

	t.Run("GetMediaInfo", func(t *testing.T) {
		info, err := client.GetMediaInfo(context.Background(), "test-media")
		require.NoError(t, err)
		assert.Equal(t, "test-media", info.ExternalID)
		assert.Equal(t, "test.txt", info.Filename)
		assert.Equal(t, "test-sha256", info.SHA256)
	})

	t.Run("GetMediaInfoBatch", func(t *testing.T) {
		externalIDs := []string{"media1", "media2"}
		mediaList, err := client.GetMediaInfoBatch(context.Background(), externalIDs)
		require.NoError(t, err)
		assert.Len(t, mediaList, 2)
		assert.Equal(t, "media1", mediaList[0].ExternalID)
		assert.Equal(t, "media2", mediaList[1].ExternalID)
	})
}

func handleGetMediaList(w http.ResponseWriter, r *http.Request) {
	response := MediaListResponse{
		ExternalIDs:     []string{"media1", "media2"},
		TotalItemsCount: 2,
		Limit:           100,
	}
	json.NewEncoder(w).Encode(response)
}

func handleUploadMedia(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // 32MB max memory
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid multipart form"}`)
		return
	}

	file, _, err := r.FormFile("media_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Missing media_file field"}`)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Failed to read file"}`)
		return
	}

	expectedContent := "test file content"
	if string(content) != expectedContent {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "File content mismatch"}`)
		return
	}

	response := struct {
		SHA256 string `json:"sha256"`
	}{
		SHA256: "test-sha256",
	}
	json.NewEncoder(w).Encode(response)
}

func handleGetMediaBinary(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test binary data"))
}

func handleGetMediaInfo(w http.ResponseWriter, r *http.Request) {
	mediaInfo := MediaInfo{
		ExternalID:  "test-media",
		Filename:    "test.txt",
		SHA256:      "test-sha256",
		CreateDate:  "2023-01-01T00:00:00Z",
		Size:        1024,
		ContentType: "text/plain",
		Description: "Test media file",
	}
	json.NewEncoder(w).Encode(mediaInfo)
}

func handleGetMediaInfoBatch(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ExternalIDs []string `json:"external_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid JSON"}`)
		return
	}

	response := struct {
		Media []MediaInfo `json:"media"`
	}{
		Media: []MediaInfo{
			{
				ExternalID:  req.ExternalIDs[0],
				Filename:    "file1.txt",
				SHA256:      "sha256-1",
				CreateDate:  "2023-01-01T00:00:00Z",
				Size:        1024,
				ContentType: "text/plain",
			},
			{
				ExternalID:  req.ExternalIDs[1],
				Filename:    "file2.jpg",
				SHA256:      "sha256-2",
				CreateDate:  "2023-01-02T00:00:00Z",
				Size:        2048,
				ContentType: "image/jpeg",
			},
		},
	}
	json.NewEncoder(w).Encode(response)
}
