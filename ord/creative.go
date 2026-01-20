package ord

import (
	"context"
	"fmt"
	"net/url"
)

const (
	CreativePayTypeCPA   = "cpa"   // Cost Per Action, цена за действие.
	CreativePayTypeCPC   = "cpc"   // Cost Per Click, цена за клик.
	CreativePayTypeCPM   = "cpm"   // Cost Per Millennium, цена за 1 000 показов.
	CreativePayTypeOther = "other" // иное.
)

const (
	CreativeFormBanner                     = "banner"                         // баннер.
	CreativeFormText                       = "text"                           // текстовый блок.
	CreativeFormAudio                      = "audio"                          // аудиозапись.
	CreativeFormVideo                      = "video"                          // видеоролик.
	CreativeFormLiveAudio                  = "live_audio"                     // аудиотрансляция в прямом эфире.
	CreativeFormLiveVideo                  = "live_video"                     // видеотрансляция в прямом эфире.
	CreativeFormTextVideoBlock             = "text_video_block"               // текстовый блок с видео
	CreativeFormTextGraphicBlock           = "text_graphic_block"             // текстово-графический блок
	CreativeFormTextAudioBlock             = "text_audio_block"               // текстовый блок с аудио
	CreativeFormTextGraphicVideoBlock      = "text_graphic_video_block"       // текстово-графический блок с видео
	CreativeFormTextAudioVideoBlock        = "text_audio_video_block"         // текстовый блок с аудио и видео
	CreativeFormTextGraphicAudioBlock      = "text_graphic_audio_block"       // текстово-графический блок с видео
	CreativeFormTextGraphicAudioVideoBlock = "text_graphic_audio_video_block" // текстово-графический блок с аудио и видео
	CreativeFormBannerHTML5                = "banner_html5"                   // HTML5-баннер
)

const (
	CreativeFlagSocial      = "social"       // социальная реклама.
	CreativeFlagNative      = "native"       // нативная реклама (только в GET, PUT не поддерживается).
	CreativeFlagSocialQuota = "social_quota" // социальная реклама по квоте.
)

type Creative struct {
	CreateDate          string    `json:"create_date,omitempty"`
	ERID                string    `json:"erid"`
	PersonExternalID    *string   `json:"person_external_id,omitempty"`
	ContractExternalID  *string   `json:"contract_external_id,omitempty"`
	ContractExternalIDs *[]string `json:"contract_external_ids,omitempty"`
	CIDs                *[]string `json:"cids,omitempty"`
	OKVEDs              *[]string `json:"okveds,omitempty"`
	KKTUs               []string  `json:"kktus"`
	Name                *string   `json:"name,omitempty"`
	Brand               *string   `json:"brand,omitempty"`
	Category            *string   `json:"category,omitempty"`
	Description         *string   `json:"description,omitempty"`
	PayType             *string   `json:"pay_type,omitempty"`
	Form                string    `json:"form"`
	Targeting           *string   `json:"targeting,omitempty"`
	TargetURLs          *[]string `json:"target_urls,omitempty"`
	Texts               *[]string `json:"texts,omitempty"`
	MediaExternalIDs    *[]string `json:"media_external_ids,omitempty"`
	MediaURLs           *[]string `json:"media_urls,omitempty"`
	Flags               *[]string `json:"flags,omitempty"`
}

// CreativeListResponse represents the response for getting a list of creatives
type CreativeListResponse struct {
	ExternalIDs     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

// CreativeERIDsListResponse represents the response for getting a list of creative ERIDs
type CreativeERIDsListResponse struct {
	ERIDs           []string `json:"erids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

// CreativeERIDExternalIDPair represents a pair of ERID and external ID
type CreativeERIDExternalIDPair struct {
	ERID       string `json:"erid"`
	ExternalID string `json:"external_id"`
}

// CreativeERIDExternalIDPairsResponse represents the response for getting pairs of ERIDs and external IDs
type CreativeERIDExternalIDPairsResponse struct {
	Items           []CreativeERIDExternalIDPair `json:"items"`
	TotalItemsCount int                          `json:"total_items_count"`
	Limit           int                          `json:"limit"`
}

// AddTextsToCreativeRequest represents the request body for adding texts to a creative
type AddTextsToCreativeRequest struct {
	Texts []string `json:"texts"`
}

// AddMediaToCreativeRequest represents the request body for adding media to a creative
type AddMediaToCreativeRequest struct {
	MediaExternalIDs []string `json:"media_external_ids"`
}

// CreateCreativeV2Request represents the request body for creating/updating a creative (v2)
type CreateCreativeV2Request struct {
	PersonExternalID   *string   `json:"person_external_id,omitempty"`
	ContractExternalID *string   `json:"contract_external_id,omitempty"`
	OKVEDs             *[]string `json:"okveds,omitempty"`
	KKTUs              []string  `json:"kktus"`
	Name               *string   `json:"name,omitempty"`
	Brand              *string   `json:"brand,omitempty"`
	Category           *string   `json:"category,omitempty"`
	Description        *string   `json:"description,omitempty"`
	PayType            *string   `json:"pay_type,omitempty"`
	Form               string    `json:"form"`
	Targeting          *string   `json:"targeting,omitempty"`
	TargetURLs         *[]string `json:"target_urls,omitempty"`
	Texts              *[]string `json:"texts,omitempty"`
	MediaExternalIDs   *[]string `json:"media_external_ids,omitempty"`
	MediaURLs          *[]string `json:"media_urls,omitempty"`
	Flags              *[]string `json:"flags,omitempty"`
}

// CreateCreativeV3Request represents the request body for creating/updating a creative (v3)
type CreateCreativeV3Request struct {
	PersonExternalID    *string   `json:"person_external_id,omitempty"`
	ContractExternalIDs *[]string `json:"contract_external_ids,omitempty"`
	CIDs                *[]string `json:"cids,omitempty"`
	KKTUs               []string  `json:"kktus"`
	Name                *string   `json:"name,omitempty"`
	Brand               *string   `json:"brand,omitempty"`
	Category            *string   `json:"category,omitempty"`
	Description         *string   `json:"description,omitempty"`
	PayType             *string   `json:"pay_type,omitempty"`
	Form                string    `json:"form"`
	Targeting           *string   `json:"targeting,omitempty"`
	TargetURLs          *[]string `json:"target_urls,omitempty"`
	Texts               *[]string `json:"texts,omitempty"`
	MediaExternalIDs    *[]string `json:"media_external_ids,omitempty"`
	MediaURLs           *[]string `json:"media_urls,omitempty"`
	Flags               *[]string `json:"flags,omitempty"`
}

// GetCreatives retrieves a list of creatives
// GET /v3/creative
func (c *Client) GetCreatives(ctx context.Context, offset, limit int) (*CreativeListResponse, error) {
	path := fmt.Sprintf("/v3/creative?offset=%d&limit=%d", offset, limit)

	var response CreativeListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get creatives: %w", err)
	}

	return &response, nil
}

// GetCreativeERIDs retrieves a list of creative ERIDs
// GET /v3/creative/list/erids
func (c *Client) GetCreativeERIDs(ctx context.Context, offset, limit int) (*CreativeERIDsListResponse, error) {
	path := fmt.Sprintf("/v3/creative/list/erids?offset=%d&limit=%d", offset, limit)

	var response CreativeERIDsListResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get creative ERIDs: %w", err)
	}

	return &response, nil
}

// GetCreativeERIDExternalIDPairs retrieves a list of pairs of ERIDs and external IDs
// GET /v3/creative/list/erid_external_ids
func (c *Client) GetCreativeERIDExternalIDPairs(ctx context.Context, offset, limit int) (*CreativeERIDExternalIDPairsResponse, error) {
	path := fmt.Sprintf("/v3/creative/list/erid_external_ids?offset=%d&limit=%d", offset, limit)

	var response CreativeERIDExternalIDPairsResponse
	if err := c.request(ctx, "GET", path, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get creative ERID/external ID pairs: %w", err)
	}

	return &response, nil
}

// CreateCreativeV2 creates or updates a creative (v2)
// PUT /v2/creative/{external_id}
func (c *Client) CreateCreativeV2(ctx context.Context, externalID string, creative CreateCreativeV2Request) error {
	path := fmt.Sprintf("/v2/creative/%s", externalID)

	if err := c.request(ctx, "PUT", path, creative, nil); err != nil {
		return fmt.Errorf("failed to create creative (v2): %w", err)
	}

	return nil
}

// GetCreativeV2 retrieves a creative by external ID (v2)
// GET /v2/creative/{external_id}
func (c *Client) GetCreativeV2(ctx context.Context, externalID string) (*Creative, error) {
	path := fmt.Sprintf("/v2/creative/%s", externalID)

	var creative Creative
	if err := c.request(ctx, "GET", path, nil, &creative); err != nil {
		return nil, fmt.Errorf("failed to get creative (v2): %w", err)
	}

	return &creative, nil
}

// GetCreativeByERIDV2 retrieves a creative by ERID (v2)
// GET /v2/creative/by_erid/{erid}
func (c *Client) GetCreativeByERIDV2(ctx context.Context, erid string) (*Creative, error) {
	path := fmt.Sprintf("/v2/creative/by_erid/%s", url.PathEscape(erid))

	var creative Creative
	if err := c.request(ctx, "GET", path, nil, &creative); err != nil {
		return nil, fmt.Errorf("failed to get creative by ERID (v2): %w", err)
	}

	return &creative, nil
}

// CreateCreativeV3 creates or updates a creative (v3)
// PUT /v3/creative/{external_id}
func (c *Client) CreateCreativeV3(ctx context.Context, externalID string, creative CreateCreativeV3Request) error {
	path := fmt.Sprintf("/v3/creative/%s", externalID)

	if err := c.request(ctx, "PUT", path, creative, nil); err != nil {
		return fmt.Errorf("failed to create creative (v3): %w", err)
	}

	return nil
}

// GetCreativeV3 retrieves a creative by external ID (v3)
// GET /v3/creative/{external_id}
func (c *Client) GetCreativeV3(ctx context.Context, externalID string) (*Creative, error) {
	path := fmt.Sprintf("/v3/creative/%s", externalID)

	var creative Creative
	if err := c.request(ctx, "GET", path, nil, &creative); err != nil {
		return nil, fmt.Errorf("failed to get creative (v3): %w", err)
	}

	return &creative, nil
}

// GetCreativeByERIDV3 retrieves a creative by ERID (v3)
// GET /v3/creative/by_erid/{erid}
func (c *Client) GetCreativeByERIDV3(ctx context.Context, erid string) (*Creative, error) {
	path := fmt.Sprintf("/v3/creative/by_erid/%s", url.PathEscape(erid))

	var creative Creative
	if err := c.request(ctx, "GET", path, nil, &creative); err != nil {
		return nil, fmt.Errorf("failed to get creative by ERID (v3): %w", err)
	}

	return &creative, nil
}

// AddTextsToCreative adds texts to a creative
// POST /v3/creative/{external_id}/add_text
func (c *Client) AddTextsToCreative(ctx context.Context, externalID string, texts []string) error {
	path := fmt.Sprintf("/v3/creative/%s/add_text", externalID)

	request := AddTextsToCreativeRequest{
		Texts: texts,
	}

	if err := c.request(ctx, "POST", path, request, nil); err != nil {
		return fmt.Errorf("failed to add texts to creative: %w", err)
	}

	return nil
}

// AddMediaToCreative adds media to a creative
// POST /v3/creative/{external_id}/add_media
func (c *Client) AddMediaToCreative(ctx context.Context, externalID string, mediaExternalIDs []string) error {
	path := fmt.Sprintf("/v3/creative/%s/add_media", externalID)

	request := AddMediaToCreativeRequest{
		MediaExternalIDs: mediaExternalIDs,
	}

	if err := c.request(ctx, "POST", path, request, nil); err != nil {
		return fmt.Errorf("failed to add media to creative: %w", err)
	}

	return nil
}
