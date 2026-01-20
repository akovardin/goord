package ord

import (
	"context"
	"fmt"
)

type Person struct {
	CreateDate       string           `json:"create_date,omitempty"`
	Name             string           `json:"name"`
	RsURL            *string          `json:"rs_url,omitempty"`
	Roles            []string         `json:"roles"`
	JuridicalDetails JuridicalDetails `json:"juridical_details"`
	LockedFields     []LockedField    `json:"locked_fields,omitempty"`
}

const (
	PersonTypePhysical         = "physical"          // физическое лицо.
	PersonTypeJuridical        = "juridical"         // юридическое лицо.
	PersonTypeIP               = "ip"                // индивидуальный предприниматель.
	PersonTypeForeignPhysical  = "foreign_physical"  // иностранное физическое лицо.
	PersonTypeForeignJuridical = "foreign_juridical" // иностранное юридическое лицо.
)

type JuridicalDetails struct {
	Type                      string  `json:"type"`
	INN                       string  `json:"inn"`
	KPP                       *string `json:"kpp,omitempty"`
	Phone                     *string `json:"phone,omitempty"`
	ForeignEpaymentMethod     *string `json:"foreign_epayment_method,omitempty"`
	ForeignRegistrationNumber *string `json:"foreign_registration_number,omitempty"`
	ForeignINN                *string `json:"foreign_inn,omitempty"`
	ForeignOKSMCountryCode    *string `json:"foreign_oksm_country_code,omitempty"`
}

type LockedField struct {
	Field   string   `json:"field"`
	Reasons []string `json:"reasons"`
}

type PersonListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

func (c *Client) GetPersons(ctx context.Context, offset, limit int) (*PersonListResponse, error) {
	path := fmt.Sprintf("/v1/person?offset=%d&limit=%d", offset, limit)

	var response PersonListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get persons: %w", err)
	}

	return &response, nil
}

func (c *Client) GetPerson(ctx context.Context, externalID string) (*Person, error) {
	path := fmt.Sprintf("/v1/person/%s", externalID)

	var person Person
	if err := c.request(ctx, "GET", path, nil, &person); err != nil {
		return nil, fmt.Errorf("failed to get person: %w", err)
	}

	return &person, nil
}

func (c *Client) CreatePerson(ctx context.Context, externalID string, person Person) error {
	path := fmt.Sprintf("/v1/person/%s", externalID)

	if err := c.request(ctx, "PUT", path, person, nil); err != nil {
		return fmt.Errorf("failed to create person: %w", err)
	}

	return nil
}
