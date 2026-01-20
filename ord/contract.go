package ord

import (
	"context"
	"fmt"
)

const (
	ContractTypeService    = "service"    // договор оказания услуг.
	ContractTypeMediation  = "mediation"  // посреднический договор. Требует заполнения поля action_type.
	ContractTypeAdditional = "additional" // дополнительное соглашение. Требует заполнения поля parent_contract_external_id.
)

const (
	ContractActionTypeDistribution = "distribution" // распространение рекламы.
	ContractActionTypeConclude     = "conclude"     // заключение договоров.
	ContractActionTypeCommercial   = "commercial"   // коммерческое представительство.
	ContractActionTypeOther        = "other"        // иное.
)

const (
	ContractSubjectTypeRepresentation  = "representation"   // представительство.
	ContractSubjectTypeOrgDistribution = "org_distribution" // организация распространения рекламы.
	ContractSubjectTypeMediation       = "mediation"        // посредничество.
	ContractSubjectTypeDistribution    = "distribution"     // распространение рекламы.
	ContractSubjectTypeOther           = "other"            // иное.
)

const (
	// vat_included — все налоги (если есть) включены в сумму договора. Обязателен для значений amount > 0.
	ContractFlagVatIncluded = "vat_included"
	// contractor_is_creatives_reporter — подрядчик обязуется вести учёт креативов.
	// Значение можно указать, только если contractor_external_id — рекламная система.
	ContractFlagContractorIsCreativesReporter = "contractor_is_creatives_reporter"
	//  деньги поступают от подрядчика (исполнителя) клиенту (заказчику).
	// Значение можно указать, только если поле type принимает значение mediation.
	ContractFlagAgentActingForPublisher = "agent_acting_for_publisher"
	// рекламный сбор в размере 3% за всю цепочку распространения рекламы оплачивает
	// исполнитель по этому договору. Значение можно указать, только если поле type
	// принимает значение mediation и подрядчик (исполнитель) не иностранный контрагент.
	// Несовместим с флагом agent_acting_for_publisher.
	ContractFlagIsChargePaidByAgent = "is_charge_paid_by_agent"
)

// Contract represents a contract (договор) in the ORD system
type Contract struct {
	CreateDate               string   `json:"create_date,omitempty"`
	Type                     string   `json:"type"`
	ClientExternalID         string   `json:"client_external_id"`
	ContractorExternalID     string   `json:"contractor_external_id"`
	ActionType               *string  `json:"action_type,omitempty"`
	SubjectType              string   `json:"subject_type"`
	Date                     string   `json:"date"`
	DateEnd                  *string  `json:"date_end,omitempty"`
	Serial                   *string  `json:"serial,omitempty"`
	Flags                    []string `json:"flags,omitempty"`
	ParentContractExternalID *string  `json:"parent_contract_external_id,omitempty"`

	// 	maxLength: 29
	// minLength: 1
	// pattern: ^\d{1,12}(\.\d{1,8})?
	// nullable: true
	// example: 500.5
	Amount *string `json:"amount,omitempty"`

	HasAdditionalContracts bool          `json:"has_additional_contracts,omitempty"`
	CID                    *string       `json:"cid,omitempty"`
	LockedFields           []LockedField `json:"locked_fields,omitempty"`
}

type ContractListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

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

func (c *Client) GetContracts(ctx context.Context, offset, limit int) (*ContractListResponse, error) {
	path := fmt.Sprintf("/v1/contract?offset=%d&limit=%d", offset, limit)

	var response ContractListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get contracts: %w", err)
	}

	return &response, nil
}

func (c *Client) GetContract(ctx context.Context, externalID string) (*Contract, error) {
	path := fmt.Sprintf("/v1/contract/%s", externalID)

	var contract Contract
	if err := c.request(ctx, "GET", path, nil, &contract); err != nil {
		return nil, fmt.Errorf("failed to get contract: %w", err)
	}

	return &contract, nil
}

func (c *Client) CreateContract(ctx context.Context, externalID string, contract CreateContractRequest) error {
	path := fmt.Sprintf("/v1/contract/%s", externalID)

	if err := c.request(ctx, "PUT", path, contract, nil); err != nil {
		return fmt.Errorf("failed to create contract: %w", err)
	}

	return nil
}

func (c *Client) RequestCID(ctx context.Context, externalID string) error {
	path := fmt.Sprintf("/v1/contract/%s/create_cid", externalID)

	if err := c.request(ctx, "POST", path, nil, nil); err != nil {
		return fmt.Errorf("failed to request CID: %w", err)
	}

	return nil
}
