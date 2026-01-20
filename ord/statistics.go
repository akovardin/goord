package ord

import (
	"context"
	"fmt"
)

// StatisticsV2Item represents a statistics item for v2 API
type StatisticsV2Item struct {
	CreativeExternalID string  `json:"creative_external_id"`
	PadExternalID      string  `json:"pad_external_id"`
	ShowsCount         uint64  `json:"shows_count"`
	InvoiceShowsCount  *uint64 `json:"invoice_shows_count,omitempty"`
	Amount             *struct {
		ExcludingVAT string `json:"excluding_vat"`
		VATRate      string `json:"vat_rate"`
		VAT          string `json:"vat"`
		IncludingVAT string `json:"including_vat"`
	} `json:"amount,omitempty"`
	AmountPerEvent   *string `json:"amount_per_event,omitempty"`
	PayType          *string `json:"pay_type,omitempty"`
	DateStartPlanned *string `json:"date_start_planned,omitempty"`
	DateEndPlanned   *string `json:"date_end_planned,omitempty"`
	DateStartActual  string  `json:"date_start_actual"`
	DateEndActual    string  `json:"date_end_actual"`
}

// StatisticsV3Item represents a statistics item for v3 API
type StatisticsV3Item struct {
	StatisticsV2Item
}

// StatisticsV2ItemsArray represents the request/response body for v2 statistics
type StatisticsV2ItemsArray struct {
	Items []StatisticsV2Item `json:"items"`
}

// StatisticsV3ItemsArray represents the request/response body for v3 statistics
type StatisticsV3ItemsArray struct {
	Items []StatisticsV3Item `json:"items"`
}

// StatisticsListResponse represents the response for getting a list of statistics
type StatisticsListResponse struct {
	Items           []StatisticsV2Item `json:"items"`
	TotalItemsCount int                `json:"total_items_count"`
	Limit           int                `json:"limit"`
}

// StatisticsExternalID represents an external ID of a statistics item
type StatisticsExternalID string

// DeleteStatisticsRequest represents the request body for deleting statistics
type DeleteStatisticsRequest struct {
	Items []struct {
		CreativeExternalID string `json:"creative_external_id"`
		PadExternalID      string `json:"pad_external_id"`
		DateStartActual    string `json:"date_start_actual"`
	} `json:"items"`
}

// CreateStatisticsV2 creates or updates statistics using v2 API
// POST /v2/statistics
func (c *Client) CreateStatisticsV2(ctx context.Context, statistics StatisticsV2ItemsArray) ([]StatisticsExternalID, error) {
	path := "/v2/statistics"

	var response struct {
		ExternalIDs []StatisticsExternalID `json:"external_ids"`
	}
	if err := c.makeRequest(ctx, "POST", path, statistics, &response); err != nil {
		return nil, fmt.Errorf("failed to create statistics v2: %w", err)
	}

	return response.ExternalIDs, nil
}

// CreateStatisticsV3 creates or updates statistics using v3 API
// POST /v3/statistics
func (c *Client) CreateStatisticsV3(ctx context.Context, statistics StatisticsV3ItemsArray) ([]StatisticsExternalID, error) {
	path := "/v3/statistics"

	var response struct {
		ExternalIDs []StatisticsExternalID `json:"external_ids"`
	}
	if err := c.makeRequest(ctx, "POST", path, statistics, &response); err != nil {
		return nil, fmt.Errorf("failed to create statistics v3: %w", err)
	}

	return response.ExternalIDs, nil
}

// GetStatisticsList retrieves a list of statistics
// GET /v2/statistics/list or /v3/statistics/list
func (c *Client) GetStatisticsList(ctx context.Context, offset, limit int) (*StatisticsListResponse, error) {
	path := fmt.Sprintf("/v3/statistics/list?offset=%d&limit=%d", offset, limit)

	var response StatisticsListResponse
	if err := c.makeRequest(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get statistics list: %w", err)
	}

	return &response, nil
}

// DeleteStatisticsV3 deletes statistics using v3 API
// POST /v3/statistics/delete
func (c *Client) DeleteStatisticsV3(ctx context.Context, deleteReq DeleteStatisticsRequest) error {
	path := "/v3/statistics/delete"

	if err := c.makeRequest(ctx, "POST", path, deleteReq, nil); err != nil {
		return fmt.Errorf("failed to delete statistics: %w", err)
	}

	return nil
}
