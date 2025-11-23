package goord

type PostV3CreativeExternalIdAddTextResponse200 struct {
	Marker   string `json:"marker"`
	Erid     string `json:"erid"`
	Messages []struct {
		Text   string `json:"text"`
		Fields []struct {
			Field  string   `json:"field"`
			Values []string `json:"values"`
		} `json:"fields"`
	} `json:"messages"`
}

type PostV3CreativeExternalIdAddTextResponse201 struct {
	Marker   string `json:"marker"`
	Erid     string `json:"erid"`
	Messages []struct {
		Text   string `json:"text"`
		Fields []struct {
			Field  string   `json:"field"`
			Values []string `json:"values"`
		} `json:"fields"`
	} `json:"messages"`
}

type PostV3CreativeExternalIdAddMediaResponse200 struct {
	Marker   string `json:"marker"`
	Erid     string `json:"erid"`
	Messages []struct {
		Text   string `json:"text"`
		Fields []struct {
			Field  string   `json:"field"`
			Values []string `json:"values"`
		} `json:"fields"`
	} `json:"messages"`
}

type PostV3CreativeExternalIdAddMediaResponse201 struct {
	Marker   string `json:"marker"`
	Erid     string `json:"erid"`
	Messages []struct {
		Text   string `json:"text"`
		Fields []struct {
			Field  string   `json:"field"`
			Values []string `json:"values"`
		} `json:"fields"`
	} `json:"messages"`
}

type GetV3CreativeListEridsResponse200 struct {
	Erids           []string `json:"erids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

type GetV3CreativeListEridExternalIdsResponse200 struct {
	Items []struct {
		Erid       string `json:"erid"`
		ExternalID string `json:"external_id"`
	} `json:"items"`
	TotalItemsCount int `json:"total_items_count"`
	Limit           int `json:"limit"`
}

type GetV3CreativeResponse200 struct {
	ExternalIds     []string `json:"external_ids"`
	TotalItemsCount int      `json:"total_items_count"`
	Limit           int      `json:"limit"`
}

type PutV3CreativeExternalIdRequest0 struct {
	Kktus               []string `json:"kktus"`
	Name                string   `json:"name"`
	Brand               string   `json:"brand"`
	Category            string   `json:"category"`
	Description         string   `json:"description"`
	PayType             string   `json:"pay_type"`
	Form                string   `json:"form"`
	Targeting           string   `json:"targeting"`
	TargetUrls          []string `json:"target_urls"`
	Texts               []string `json:"texts"`
	MediaExternalIds    []string `json:"media_external_ids"`
	Flags               []string `json:"flags"`
	ContractExternalIds []string `json:"contract_external_ids"`
}

type PutV3CreativeExternalIdRequest1 struct {
	Kktus               []string `json:"kktus"`
	Name                string   `json:"name"`
	Brand               string   `json:"brand"`
	Category            string   `json:"category"`
	Description         string   `json:"description"`
	PayType             string   `json:"pay_type"`
	Form                string   `json:"form"`
	Targeting           string   `json:"targeting"`
	TargetUrls          []string `json:"target_urls"`
	Texts               []string `json:"texts"`
	MediaExternalIds    []string `json:"media_external_ids"`
	Flags               []string `json:"flags"`
	ContractExternalIds []string `json:"contract_external_ids"`
}

type PutV3CreativeExternalIdRequest2 struct {
	Kktus               []string `json:"kktus"`
	Name                string   `json:"name"`
	Brand               string   `json:"brand"`
	Category            string   `json:"category"`
	Description         string   `json:"description"`
	PayType             string   `json:"pay_type"`
	Form                string   `json:"form"`
	Targeting           string   `json:"targeting"`
	TargetUrls          []string `json:"target_urls"`
	Texts               []string `json:"texts"`
	MediaExternalIds    []string `json:"media_external_ids"`
	Flags               []string `json:"flags"`
	ContractExternalIds []string `json:"contract_external_ids"`
}

type GetV3CreativeExternalIdResponse2000 struct {
	Erid                string   `json:"erid"`
	Kktus               []string `json:"kktus"`
	Name                string   `json:"name"`
	Brand               string   `json:"brand"`
	Category            string   `json:"category"`
	Description         string   `json:"description"`
	PayType             string   `json:"pay_type"`
	Form                string   `json:"form"`
	Targeting           string   `json:"targeting"`
	TargetUrls          []string `json:"target_urls"`
	Texts               []string `json:"texts"`
	MediaExternalIds    []string `json:"media_external_ids"`
	Flags               []string `json:"flags"`
	ContractExternalIds []string `json:"contract_external_ids"`
}

type GetV3CreativeExternalIdResponse2001 struct {
	Erid                string   `json:"erid"`
	Kktus               []string `json:"kktus"`
	Name                string   `json:"name"`
	Brand               string   `json:"brand"`
	Category            string   `json:"category"`
	Description         string   `json:"description"`
	PayType             string   `json:"pay_type"`
	Form                string   `json:"form"`
	Targeting           string   `json:"targeting"`
	TargetUrls          []string `json:"target_urls"`
	Texts               []string `json:"texts"`
	MediaExternalIds    []string `json:"media_external_ids"`
	Flags               []string `json:"flags"`
	ContractExternalIds []string `json:"contract_external_ids"`
}

type GetV3CreativeExternalIdResponse2002 struct {
	Erid                string   `json:"erid"`
	Kktus               []string `json:"kktus"`
	Name                string   `json:"name"`
	Brand               string   `json:"brand"`
	Category            string   `json:"category"`
	Description         string   `json:"description"`
	PayType             string   `json:"pay_type"`
	Form                string   `json:"form"`
	Targeting           string   `json:"targeting"`
	TargetUrls          []string `json:"target_urls"`
	Texts               []string `json:"texts"`
	MediaExternalIds    []string `json:"media_external_ids"`
	Flags               []string `json:"flags"`
	ContractExternalIds []string `json:"contract_external_ids"`
}
