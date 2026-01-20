package ord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// MediaInfo represents information about a media file
type MediaInfo struct {
	ExternalID  string `json:"external_id"`
	Filename    string `json:"filename"`
	SHA256      string `json:"sha256"`
	CreateDate  string `json:"create_date"`
	Size        int64  `json:"size"`
	ContentType string `json:"content_type"`
	Description string `json:"description,omitempty"`
}

// MediaListResponse represents the response for getting a list of media files
type MediaListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

// GetMediaList retrieves a list of media files
// GET /v1/media
func (c *Client) GetMediaList(ctx context.Context, offset, limit int) (*MediaListResponse, error) {
	path := fmt.Sprintf("/v1/media?offset=%d&limit=%d", offset, limit)

	var response MediaListResponse
	if err := c.makeRequest(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get media list: %w", err)
	}

	return &response, nil
}

// UploadMedia uploads a media file
// PUT /v1/media/{external_id}
func (c *Client) UploadMedia(ctx context.Context, externalID string, filename string, fileReader io.Reader) (*string, error) {
	path := fmt.Sprintf("/v1/media/%s", url.PathEscape(externalID))

	// Create a buffer to write our multipart form
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Create the file field
	fw, err := w.CreateFormFile("media_file", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	// Copy the file data to the field
	if _, err = io.Copy(fw, fileReader); err != nil {
		return nil, fmt.Errorf("failed to copy file data: %w", err)
	}

	// Close the multipart writer
	w.Close()

	// Construct the full URL
	url := c.baseURL + path

	// Create request with body
	req, err := http.NewRequestWithContext(ctx, "PUT", url, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Set authorization header
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	// Perform the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse the response to get the SHA256
	var result struct {
		SHA256 string `json:"sha256"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result.SHA256, nil
}

// GetMediaBinary retrieves the binary representation of a media file
// GET /v1/media/{external_id}
func (c *Client) GetMediaBinary(ctx context.Context, externalID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/media/%s", url.PathEscape(externalID))

	// Construct the full URL
	url := c.baseURL + path

	// Create request without body
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set authorization header
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	// Perform the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// GetMediaInfo retrieves data about a media file
// GET /v1/media/{external_id}/info
func (c *Client) GetMediaInfo(ctx context.Context, externalID string) (*MediaInfo, error) {
	path := fmt.Sprintf("/v1/media/%s/info", url.PathEscape(externalID))

	var mediaInfo MediaInfo
	if err := c.makeRequest(ctx, "GET", path, nil, &mediaInfo); err != nil {
		return nil, fmt.Errorf("failed to get media info: %w", err)
	}

	return &mediaInfo, nil
}

// GetMediaInfoBatch retrieves data about multiple media files
// POST /v1/get_media_info
func (c *Client) GetMediaInfoBatch(ctx context.Context, externalIDs []string) ([]MediaInfo, error) {
	path := "/v1/get_media_info"

	request := struct {
		ExternalIDs []string `json:"external_ids"`
	}{
		ExternalIDs: externalIDs,
	}

	var response struct {
		Media []MediaInfo `json:"media"`
	}
	if err := c.makeRequest(ctx, "POST", path, request, &response); err != nil {
		return nil, fmt.Errorf("failed to get media info batch: %w", err)
	}

	return response.Media, nil
}
