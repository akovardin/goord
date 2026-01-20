package ord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DictionaryService handles communication with the dictionary related methods of the ORD API
type DictionaryService struct {
	client *Client
}

// KKTUItem represents a KKTU code with its description
type KKTUItem struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// KKTUResponse represents the response for KKTU codes
type KKTUResponse struct {
	TotalItemsCount int        `json:"total_items_count"`
	Limit           int        `json:"limit"`
	Items           []KKTUItem `json:"items"`
}

// ERIRMessageItem represents an ERIR message with its translation
type ERIRMessageItem struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// ERIRMessageResponse represents the response for ERIR messages
type ERIRMessageResponse struct {
	Items []ERIRMessageItem `json:"items"`
}

// GetKKTUCodes retrieves KKTU codes
func (s *DictionaryService) GetKKTUCodes(ctx context.Context, search, lang string, offset, limit int, codes []string) (*KKTUResponse, error) {
	params := url.Values{}
	if search != "" {
		params.Set("search", search)
	}
	if lang != "" {
		params.Set("lang", lang)
	}
	if offset > 0 {
		params.Set("offset", fmt.Sprintf("%d", offset))
	}
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if len(codes) > 0 {
		params.Set("codes", fmt.Sprintf("%v", codes))
	}

	path := "/v1/dict/kktu"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var response KKTUResponse
	err := s.client.makeRequest(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetERIRMessage retrieves translation for a single ERIR message
func (s *DictionaryService) GetERIRMessage(ctx context.Context, lang, message string) (*ERIRMessageResponse, error) {
	params := url.Values{}
	if lang != "" {
		params.Set("lang", lang)
	}
	if message != "" {
		params.Set("message", message)
	}

	path := "/v1/dict/erir_message"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var response ERIRMessageResponse
	err := s.client.makeRequest(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// PostERIRMessages retrieves translations for multiple ERIR messages
func (s *DictionaryService) PostERIRMessages(ctx context.Context, lang string, messages []string) (*ERIRMessageResponse, error) {
	type request struct {
		Lang     string   `json:"lang,omitempty"`
		Messages []string `json:"messages"`
	}

	req := request{
		Lang:     lang,
		Messages: messages,
	}

	var response ERIRMessageResponse
	err := s.client.makeRequest(ctx, http.MethodPost, "/v1/dict/erir_message", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
