package ord

import (
	"context"
	"fmt"
)

// Pad represents an advertising venue (рекламная площадка) in the ORD system
type Pad struct {
	CreateDate       string  `json:"create_date,omitempty"`
	PersonExternalID string  `json:"person_external_id"`
	IsOwner          bool    `json:"is_owner"`
	Type             string  `json:"type"`
	Name             string  `json:"name"`
	URL              *string `json:"url,omitempty"`
}

// PadListResponse represents the response for getting a list of pads
type PadListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

// GetPads retrieves a list of advertising venues (рекламных площадок)
// GET /v1/pad
func (c *Client) GetPads(ctx context.Context, offset, limit int, personExternalID string) (*PadListResponse, error) {
	path := fmt.Sprintf("/v1/pad?offset=%d&limit=%d", offset, limit)
	if personExternalID != "" {
		path = fmt.Sprintf("%s&person_external_id=%s", path, personExternalID)
	}

	var response PadListResponse
	if err := c.makeRequest(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get pads: %w", err)
	}

	return &response, nil
}

// GetRestrictedPads retrieves a list of restricted pad URLs
// GET /v1/pad/info/restricted
func (c *Client) GetRestrictedPads(ctx context.Context) ([]string, error) {
	path := "/v1/pad/info/restricted"

	var response struct {
		URLs []string `json:"urls"`
	}
	if err := c.makeRequest(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get restricted pads: %w", err)
	}

	return response.URLs, nil
}

// GetPad retrieves an advertising venue by external ID
// GET /v1/pad/{external_id}
func (c *Client) GetPad(ctx context.Context, externalID string) (*Pad, error) {
	path := fmt.Sprintf("/v1/pad/%s", externalID)

	var pad Pad
	if err := c.makeRequest(ctx, "GET", path, nil, &pad); err != nil {
		return nil, fmt.Errorf("failed to get pad: %w", err)
	}

	return &pad, nil
}

// CreatePad creates or updates an advertising venue
// PUT /v1/pad/{external_id}
func (c *Client) CreatePad(ctx context.Context, externalID string, pad Pad) error {
	path := fmt.Sprintf("/v1/pad/%s", externalID)

	if err := c.makeRequest(ctx, "PUT", path, pad, nil); err != nil {
		return fmt.Errorf("failed to create pad: %w", err)
	}

	return nil
}
