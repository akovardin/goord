package ord

import (
	"context"
	"fmt"
)

const (
	InvoiceClientRoleTypeAdvertiser = "advertiser" // рекламодатель.
	InvoiceClientRoleTypeAgency     = "agency"     // рекламное агентство.
	InvoiceClientRoleTypeOrs        = "ors"        // оператор рекламной системы.
	InvoiceClientRoleTypePublisher  = "publisher"  // издатель, рекламораспространитель.
	InvoiceClientRoleTypeMediator   = "mediator"   // посредник
)

type Invoice struct {
	ContractExternalID      string        `json:"contract_external_id"`
	OrderContractExternalID *string       `json:"order_contract_external_id,omitempty"`
	Date                    string        `json:"date"`
	Serial                  *string       `json:"serial,omitempty"`
	DateStart               string        `json:"date_start"`
	DateEnd                 string        `json:"date_end"`
	Amount                  InvoiceAmount `json:"amount"`
	ClientRole              string        `json:"client_role"`
	ContractorRole          string        `json:"contractor_role"`
	Flags                   []string      `json:"flags,omitempty"`
	Items                   []InvoiceItem `json:"items,omitempty"`
	Status                  *string       `json:"status,omitempty"`
	ErirTaxStatus           *string       `json:"erir_tax_status,omitempty"`
}

type InvoiceAmount struct {
	Services   InvoiceAmountGroup `json:"services"`
	Commission *InvoiceCommission `json:"commission,omitempty"`
}

type InvoiceAmountGroup struct {
	ExcludingVat string `json:"excluding_vat"`
	VatRate      string `json:"vat_rate"`
	Vat          string `json:"vat"`
	IncludingVat string `json:"including_vat"`
}

type InvoiceCommission struct {
	Serial *string            `json:"serial,omitempty"`
	Date   *string            `json:"date,omitempty"`
	Amount InvoiceAmountGroup `json:"amount"`
}

type InvoiceItem struct {
	ContractExternalID *string            `json:"contract_external_id,omitempty"`
	Cid                *string            `json:"cid,omitempty"`
	Amount             InvoiceAmountGroup `json:"amount"`
	Creatives          []InvoiceCreative  `json:"creatives,omitempty"`
}

type InvoiceCreative struct {
	CreativeExternalID string                    `json:"creative_external_id"`
	Platforms          []InvoiceCreativePlatform `json:"platforms,omitempty"`
}

type InvoiceCreativePlatform struct {
	PadExternalID     string             `json:"pad_external_id"`
	ShowsCount        int64              `json:"shows_count"`
	InvoiceShowsCount int64              `json:"invoice_shows_count"`
	Amount            InvoiceAmountGroup `json:"amount"`
	AmountPerEvent    *string            `json:"amount_per_event,omitempty"`
	DateStartPlanned  string             `json:"date_start_planned"`
	DateEndPlanned    string             `json:"date_end_planned"`
	DateStartActual   string             `json:"date_start_actual"`
	DateEndActual     string             `json:"date_end_actual"`
	PayType           string             `json:"pay_type"`
}

type InvoiceListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

func (c *Client) GetInvoices(ctx context.Context, offset, limit int) (*InvoiceListResponse, error) {
	path := fmt.Sprintf("/v1/invoice?offset=%d&limit=%d", offset, limit)

	var response InvoiceListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get invoices: %w", err)
	}

	return &response, nil
}

func (c *Client) GetInvoice(ctx context.Context, externalID string) (*Invoice, error) {
	path := fmt.Sprintf("/v4/invoice/%s", externalID)

	var invoice Invoice
	if err := c.request(ctx, "GET", path, nil, &invoice); err != nil {
		return nil, fmt.Errorf("failed to get invoice: %w", err)
	}

	return &invoice, nil
}

func (c *Client) CreateInvoiceHeader(ctx context.Context, externalID string, invoice Invoice) error {
	path := fmt.Sprintf("/v4/invoice/%s/header", externalID)

	if err := c.request(ctx, "PUT", path, invoice, nil); err != nil {
		return fmt.Errorf("failed to create invoice header: %w", err)
	}

	return nil
}

func (c *Client) AddContractsToInvoice(ctx context.Context, externalID string, items []InvoiceItem) error {
	path := fmt.Sprintf("/v4/invoice/%s/items", externalID)

	requestBody := map[string]interface{}{
		"items": items,
	}

	if err := c.request(ctx, "PATCH", path, requestBody, nil); err != nil {
		return fmt.Errorf("failed to add contracts to invoice: %w", err)
	}

	return nil
}

func (c *Client) DeleteInvoice(ctx context.Context, externalID string) error {
	path := fmt.Sprintf("/v4/invoice/%s", externalID)

	if err := c.request(ctx, "DELETE", path, nil, nil); err != nil {
		return fmt.Errorf("failed to delete invoice: %w", err)
	}

	return nil
}

func (c *Client) SendInvoiceToErir(ctx context.Context, externalID string) error {
	path := fmt.Sprintf("/v4/invoice/%s/ready", externalID)

	if err := c.request(ctx, "POST", path, nil, nil); err != nil {
		return fmt.Errorf("failed to send invoice to ERIR: %w", err)
	}

	return nil
}

func (c *Client) DeleteContractsFromInvoice(ctx context.Context, externalID string, deleteInfo interface{}) error {
	path := fmt.Sprintf("/v4/invoice/%s/delete", externalID)

	if err := c.request(ctx, "POST", path, deleteInfo, nil); err != nil {
		return fmt.Errorf("failed to delete contracts from invoice: %w", err)
	}

	return nil
}

func (c *Client) CreateWholeInvoice(ctx context.Context, externalID string, invoice Invoice) error {
	path := fmt.Sprintf("/v4/invoice/%s", externalID)

	if err := c.request(ctx, "PUT", path, invoice, nil); err != nil {
		return fmt.Errorf("failed to create whole invoice: %w", err)
	}

	return nil
}
