package ord

import (
	"context"
	"fmt"
)

type CID struct {
	CID                             string  `json:"cid"`
	Name                            string  `json:"name"`
	ErirStatus                      string  `json:"erir_status"`
	ClientINN                       *string `json:"client_inn,omitempty"`
	ClientPhone                     *string `json:"client_phone,omitempty"`
	ClientForeignEpaymentMethod     *string `json:"client_foreign_epayment_method,omitempty"`
	ClientForeignRegistrationNumber *string `json:"client_foreign_registration_number,omitempty"`
	ClientForeignINN                *string `json:"client_foreign_inn,omitempty"`
}

type CIDListResponse struct {
	CIDs            []string `json:"cids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

func (c *Client) GetCIDList(ctx context.Context, offset, limit int) (*CIDListResponse, error) {
	path := fmt.Sprintf("/v1/cid?offset=%d&limit=%d", offset, limit)

	var response CIDListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get CID list: %w", err)
	}

	return &response, nil
}

func (c *Client) GetCID(ctx context.Context, cidValue string) (*CID, error) {
	path := fmt.Sprintf("/v1/cid/%s", cidValue)

	var cid CID
	if err := c.request(ctx, "GET", path, nil, &cid); err != nil {
		return nil, fmt.Errorf("failed to get CID: %w", err)
	}

	return &cid, nil
}

func (c *Client) CreateCID(ctx context.Context, cidValue string, cid CID) error {
	path := fmt.Sprintf("/v1/cid/%s", cidValue)

	if err := c.request(ctx, "PUT", path, cid, nil); err != nil {
		return fmt.Errorf("failed to create CID: %w", err)
	}

	return nil
}
