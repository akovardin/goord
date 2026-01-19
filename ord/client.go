package ord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents the ORD VK API client
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// Config represents the client configuration
type Config struct {
	BaseURL string
	Token   string
	Timeout time.Duration
}

// NewClient creates a new ORD VK API client
func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.ord.vk.com"
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &Client{
		baseURL: config.BaseURL,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
		token: config.Token,
	}
}

// SetToken sets the authentication token
func (c *Client) SetToken(token string) {
	c.token = token
}

// makeRequest performs an HTTP request to the ORD VK API
func (c *Client) makeRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	// Construct the full URL
	url := c.baseURL + path

	// Create the request
	var req *http.Request
	var err error

	if body != nil {
		// Marshal the body to JSON
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		// Create request with body
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		// Set content type
		req.Header.Set("Content-Type", "application/json")
	} else {
		// Create request without body
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
	}

	// Set authorization header
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	// Perform the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Unmarshal the response if a result struct is provided
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
