package ord

import (
	"context"
	"fmt"
)

const (
	PadTypeWeb       = "web"        // веб-страница, включая мобильные версии сайтов или профили социальной сети.
	PadTypeMobileApp = "mobile_app" // приложение.
	PadTypeHbbTV     = "hbbtv"      // приложение HbbTV
	PadTypeSmartTV   = "smarttv"    // приложение SmartTV
)

type Pad struct {
	CreateDate       string  `json:"create_date,omitempty"`
	PersonExternalID string  `json:"person_external_id"`
	IsOwner          bool    `json:"is_owner"`
	Type             string  `json:"type"`
	Name             string  `json:"name"`
	URL              *string `json:"url,omitempty"`
}

type PadListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

func (c *Client) GetPads(ctx context.Context, offset, limit int, personExternalID string) (*PadListResponse, error) {
	path := fmt.Sprintf("/v1/pad?offset=%d&limit=%d", offset, limit)
	if personExternalID != "" {
		path = fmt.Sprintf("%s&person_external_id=%s", path, personExternalID)
	}

	var response PadListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get pads: %w", err)
	}

	return &response, nil
}

func (c *Client) GetRestrictedPads(ctx context.Context) ([]string, error) {
	path := "/v1/pad/info/restricted"

	var response struct {
		URLs []string `json:"urls"`
	}
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get restricted pads: %w", err)
	}

	return response.URLs, nil
}

func (c *Client) GetPad(ctx context.Context, externalID string) (*Pad, error) {
	path := fmt.Sprintf("/v1/pad/%s", externalID)

	var pad Pad
	if err := c.request(ctx, "GET", path, nil, &pad); err != nil {
		return nil, fmt.Errorf("failed to get pad: %w", err)
	}

	return &pad, nil
}

func (c *Client) CreatePad(ctx context.Context, externalID string, pad Pad) error {
	path := fmt.Sprintf("/v1/pad/%s", externalID)

	if err := c.request(ctx, "PUT", path, pad, nil); err != nil {
		return fmt.Errorf("failed to create pad: %w", err)
	}

	return nil
}
