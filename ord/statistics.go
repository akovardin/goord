package ord

import (
	"context"
	"fmt"
)

const (
	StatisticsPayTypeCPA   = "cpa"   // Cost Per Action, цена за действие.
	StatisticsPayTypeCPC   = "cpc"   // Cost Per Click, цена за клик.
	StatisticsPayTypeCPM   = "cpm"   // Cost Per Millennium, цена за 1 000 показов.
	PStatisticsayTypeOther = "other" // иное.
)

type StatisticsV2Item struct {
	CreativeExternalID string            `json:"creative_external_id"`
	PadExternalID      string            `json:"pad_external_id"`
	ShowsCount         uint64            `json:"shows_count"`
	InvoiceShowsCount  *uint64           `json:"invoice_shows_count,omitempty"`
	Amount             *StatisticsAmount `json:"amount,omitempty"`
	AmountPerEvent     *string           `json:"amount_per_event,omitempty"`
	PayType            *string           `json:"pay_type,omitempty"`
	DateStartPlanned   *string           `json:"date_start_planned,omitempty"`
	DateEndPlanned     *string           `json:"date_end_planned,omitempty"`
	DateStartActual    string            `json:"date_start_actual"`
	DateEndActual      string            `json:"date_end_actual"`
}

type StatisticsAmount struct {
	ExcludingVAT string `json:"excluding_vat"`
	VATRate      string `json:"vat_rate"`
	VAT          string `json:"vat"`
	IncludingVAT string `json:"including_vat"`
}

type StatisticsV3Item struct {
	StatisticsV2Item
}

type StatisticsV2ItemsArray struct {
	Items []StatisticsV2Item `json:"items"`
}

type StatisticsV3ItemsArray struct {
	Items []StatisticsV3Item `json:"items"`
}

type StatisticsListResponse struct {
	Items           []StatisticsV2Item `json:"items"`
	TotalItemsCount int                `json:"total_items_count"`
	Limit           int                `json:"limit"`
}

type StatisticsExternalID string

type DeleteStatisticsRequest struct {
	Items []struct {
		CreativeExternalID string `json:"creative_external_id"`
		PadExternalID      string `json:"pad_external_id"`
		DateStartActual    string `json:"date_start_actual"`
	} `json:"items"`
}

func (c *Client) CreateStatisticsV2(ctx context.Context, statistics StatisticsV2ItemsArray) ([]StatisticsExternalID, error) {
	path := "/v2/statistics"

	var response struct {
		ExternalIDs []StatisticsExternalID `json:"external_ids"`
	}
	if err := c.request(ctx, "POST", path, statistics, &response); err != nil {
		return nil, fmt.Errorf("failed to create statistics v2: %w", err)
	}

	return response.ExternalIDs, nil
}

func (c *Client) CreateStatisticsV3(ctx context.Context, statistics StatisticsV3ItemsArray) ([]StatisticsExternalID, error) {
	path := "/v3/statistics"

	var response struct {
		ExternalIDs []StatisticsExternalID `json:"external_ids"`
	}
	if err := c.request(ctx, "POST", path, statistics, &response); err != nil {
		return nil, fmt.Errorf("failed to create statistics v3: %w", err)
	}

	return response.ExternalIDs, nil
}

func (c *Client) GetStatisticsList(ctx context.Context, offset, limit int) (*StatisticsListResponse, error) {
	path := fmt.Sprintf("/v3/statistics/list?offset=%d&limit=%d", offset, limit)

	var response StatisticsListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get statistics list: %w", err)
	}

	return &response, nil
}

func (c *Client) DeleteStatisticsV3(ctx context.Context, deleteReq DeleteStatisticsRequest) error {
	path := "/v3/statistics/delete"

	if err := c.request(ctx, "POST", path, deleteReq, nil); err != nil {
		return fmt.Errorf("failed to delete statistics: %w", err)
	}

	return nil
}
