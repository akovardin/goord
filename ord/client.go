package ord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	base  string
	http  *http.Client
	token string
}

func NewClient(options ...Option) (*Client, error) {
	cl := &Client{
		base: "https://api.ord.vk.com",
		http: http.DefaultClient,
	}

	for _, o := range options {
		if err := o(cl); err != nil {
			return nil, err
		}
	}

	return cl, nil
}

func (c *Client) SetToken(token string) {
	c.token = token
}

// request performs an HTTP request to the ORD VK API
func (c *Client) request(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	url := c.base + path

	var req *http.Request
	var err error

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("error on close body", err)
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
