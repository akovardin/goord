package ord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

type MediaInfo struct {
	ExternalID  string `json:"external_id"`
	Filename    string `json:"filename"`
	SHA256      string `json:"sha256"`
	CreateDate  string `json:"create_date"`
	Size        int64  `json:"size"`
	ContentType string `json:"content_type"`
	Description string `json:"description,omitempty"`
}

type MediaListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

func (c *Client) GetMediaList(ctx context.Context, offset, limit int) (*MediaListResponse, error) {
	path := fmt.Sprintf("/v1/media?offset=%d&limit=%d", offset, limit)

	var response MediaListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get media list: %w", err)
	}

	return &response, nil
}

func (c *Client) UploadMedia(ctx context.Context, externalID string, filename string, fileReader io.Reader) (*string, error) {
	path := fmt.Sprintf("/v1/media/%s", url.PathEscape(externalID))

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("media_file", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err = io.Copy(fw, fileReader); err != nil {
		return nil, fmt.Errorf("failed to copy file data: %w", err)
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	url := c.base + path

	req, err := http.NewRequestWithContext(ctx, "PUT", url, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("failed to close response body", err)
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		SHA256 string `json:"sha256"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result.SHA256, nil
}

func (c *Client) GetMediaBinary(ctx context.Context, externalID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/media/%s", url.PathEscape(externalID))

	url := c.base + path

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("failed to close response body", err)
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *Client) GetMediaInfo(ctx context.Context, externalID string) (*MediaInfo, error) {
	path := fmt.Sprintf("/v1/media/%s/info", url.PathEscape(externalID))

	var mediaInfo MediaInfo
	if err := c.request(ctx, "GET", path, nil, &mediaInfo); err != nil {
		return nil, fmt.Errorf("failed to get media info: %w", err)
	}

	return &mediaInfo, nil
}

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
	if err := c.request(ctx, "POST", path, request, &response); err != nil {
		return nil, fmt.Errorf("failed to get media info batch: %w", err)
	}

	return response.Media, nil
}
