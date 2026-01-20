package ord

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type KKTUItem struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type KKTUResponse struct {
	TotalItemsCount int        `json:"total_items_count"`
	Limit           int        `json:"limit"`
	Items           []KKTUItem `json:"items"`
}

type ERIRMessageItem struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

type ERIRMessageResponse struct {
	Items []ERIRMessageItem `json:"items"`
}

func (s *Client) GetKKTUCodes(ctx context.Context, search, lang string, offset, limit int, codes []string) (*KKTUResponse, error) {
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
	err := s.request(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *Client) GetERIRMessage(ctx context.Context, lang, message string) (*ERIRMessageResponse, error) {
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
	err := s.request(ctx, http.MethodGet, path, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *Client) PostERIRMessages(ctx context.Context, lang string, messages []string) (*ERIRMessageResponse, error) {
	type request struct {
		Lang     string   `json:"lang,omitempty"`
		Messages []string `json:"messages"`
	}

	req := request{
		Lang:     lang,
		Messages: messages,
	}

	var response ERIRMessageResponse
	err := s.request(ctx, http.MethodPost, "/v1/dict/erir_message", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
