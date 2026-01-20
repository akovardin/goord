package ord

import (
	"context"
	"fmt"
	"net/url"
)

// ErirStatusEntity represents the status of processing an advertising object
type ErirStatusEntity struct {
	ErirStatus      string   `json:"erir_status"`
	UpdatedByUserTs string   `json:"updated_by_user_ts"`
	FinalizedTs     *string  `json:"finalized_ts,omitempty"`
	Messages        []string `json:"messages,omitempty"`
}

// ErirStatusEntityItem represents an item in the ErirStatusEntities response
type ErirStatusEntityItem struct {
	DataType        string   `json:"data_type"`
	ExternalID      string   `json:"external_id"`
	Name            string   `json:"name"`
	ErirTaxStatus   string   `json:"erir_tax_status"`
	ErirStatus      string   `json:"erir_status"`
	UpdatedByUserTs string   `json:"updated_by_user_ts"`
	FinalizedTs     *string  `json:"finalized_ts,omitempty"`
	Messages        []string `json:"messages,omitempty"`
}

// ErirStatusEntities represents the response for getting statuses of advertising objects
type ErirStatusEntities struct {
	TotalItemsCount int                    `json:"total_items_count"`
	Limit           int                    `json:"limit"`
	LimitPerEntity  int                    `json:"limit_per_entity"`
	Items           []ErirStatusEntityItem `json:"items"`
}

// GetObjectProcessingStatus retrieves the processing status of an advertising object
// GET /v1/{data_type}/{external_id}/erir_status
func (c *Client) GetObjectProcessingStatus(ctx context.Context, dataType, externalID string) (*ErirStatusEntity, error) {
	path := fmt.Sprintf("/v1/%s/%s/erir_status", dataType, externalID)

	var status ErirStatusEntity
	if err := c.makeRequest(ctx, "GET", path, nil, &status); err != nil {
		return nil, fmt.Errorf("failed to get object processing status: %w", err)
	}

	return &status, nil
}

// GetAdObjectProcessingStatus retrieves the processing statuses of advertising objects
// GET /v1/erir_statuses
func (c *Client) GetAdObjectProcessingStatus(ctx context.Context, dataType, erirStatus string, offset, limit, limitPerEntity int, externalIDs []string) (*ErirStatusEntities, error) {
	// Build query parameters
	params := url.Values{}
	if dataType != "" {
		params.Set("data_type", dataType)
	}
	if erirStatus != "" {
		params.Set("erir_status", erirStatus)
	}
	params.Set("offset", fmt.Sprintf("%d", offset))
	params.Set("limit", fmt.Sprintf("%d", limit))
	params.Set("limit_per_entity", fmt.Sprintf("%d", limitPerEntity))

	for _, id := range externalIDs {
		params.Add("external_id", id)
	}

	path := "/v1/erir_statuses?" + params.Encode()

	var statuses ErirStatusEntities
	if err := c.makeRequest(ctx, "GET", path, nil, &statuses); err != nil {
		return nil, fmt.Errorf("failed to get ad object processing statuses: %w", err)
	}

	return &statuses, nil
}

// PostAdObjectProcessingStatus retrieves the processing statuses of advertising objects using POST method
// POST /v1/erir_statuses
type PostAdObjectProcessingStatusRequest struct {
	DataType       string   `json:"data_type,omitempty"`
	ErirStatus     string   `json:"erir_status,omitempty"`
	ExternalID     []string `json:"external_id,omitempty"`
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	LimitPerEntity int      `json:"limit_per_entity,omitempty"`
}

func (c *Client) PostAdObjectProcessingStatus(ctx context.Context, request PostAdObjectProcessingStatusRequest) (*ErirStatusEntities, error) {
	path := "/v1/erir_statuses"

	var statuses ErirStatusEntities
	if err := c.makeRequest(ctx, "POST", path, request, &statuses); err != nil {
		return nil, fmt.Errorf("failed to post ad object processing statuses: %w", err)
	}

	return &statuses, nil
}
