package ord

import (
	"context"
	"fmt"
)

// Contract represents a contract (договор) in the ORD system
type Contract struct {
	CreateDate               string        `json:"create_date,omitempty"`
	Type                     string        `json:"type"`
	ClientExternalID         string        `json:"client_external_id"`
	ContractorExternalID     string        `json:"contractor_external_id"`
	ActionType               *string       `json:"action_type,omitempty"`
	SubjectType              string        `json:"subject_type"`
	Date                     string        `json:"date"`
	DateEnd                  *string       `json:"date_end,omitempty"`
	Serial                   *string       `json:"serial,omitempty"`
	Flags                    []string      `json:"flags,omitempty"`
	ParentContractExternalID *string       `json:"parent_contract_external_id,omitempty"`
	Amount                   *string       `json:"amount,omitempty"`
	HasAdditionalContracts   bool          `json:"has_additional_contracts,omitempty"`
	CID                      *string       `json:"cid,omitempty"`
	LockedFields             []LockedField `json:"locked_fields,omitempty"`
}

// ContractListResponse represents the response for getting a list of contracts
type ContractListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

// CreateContractRequest represents the request body for creating/updating a contract
type CreateContractRequest struct {
	Type                     string   `json:"type"`
	ClientExternalID         string   `json:"client_external_id"`
	ContractorExternalID     string   `json:"contractor_external_id"`
	Date                     string   `json:"date"`
	DateEnd                  *string  `json:"date_end,omitempty"`
	Serial                   *string  `json:"serial,omitempty"`
	ActionType               *string  `json:"action_type,omitempty"`
	SubjectType              string   `json:"subject_type"`
	Flags                    []string `json:"flags,omitempty"`
	ParentContractExternalID *string  `json:"parent_contract_external_id,omitempty"`
	Amount                   *string  `json:"amount,omitempty"`
}

// GetContracts retrieves a list of contracts (договоров)
// GET /v1/contract
func (c *Client) GetContracts(ctx context.Context, offset, limit int) (*ContractListResponse, error) {
	path := fmt.Sprintf("/v1/contract?offset=%d&limit=%d", offset, limit)

	var response ContractListResponse
	if err := c.makeRequest(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get contracts: %w", err)
	}

	return &response, nil
}

// GetContract retrieves a contract by external ID
// GET /v1/contract/{external_id}
func (c *Client) GetContract(ctx context.Context, externalID string) (*Contract, error) {
	path := fmt.Sprintf("/v1/contract/%s", externalID)

	var contract Contract
	if err := c.makeRequest(ctx, "GET", path, nil, &contract); err != nil {
		return nil, fmt.Errorf("failed to get contract: %w", err)
	}

	return &contract, nil
}

// CreateContract creates or updates a contract
// PUT /v1/contract/{external_id}
func (c *Client) CreateContract(ctx context.Context, externalID string, contract CreateContractRequest) error {
	path := fmt.Sprintf("/v1/contract/%s", externalID)

	if err := c.makeRequest(ctx, "PUT", path, contract, nil); err != nil {
		return fmt.Errorf("failed to create contract: %w", err)
	}

	return nil
}

// RequestCID requests a CID for a contract
// POST /v1/contract/{external_id}/create_cid
func (c *Client) RequestCID(ctx context.Context, externalID string) error {
	path := fmt.Sprintf("/v1/contract/%s/create_cid", externalID)

	if err := c.makeRequest(ctx, "POST", path, nil, nil); err != nil {
		return fmt.Errorf("failed to request CID: %w", err)
	}

	return nil
}
