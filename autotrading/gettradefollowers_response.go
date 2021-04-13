package autotrading

type GetTradeFollowersResponse struct {
	ClientLegalAssetTypes         []string `json:"ClientLegalAssetTypes"`
	ClientSuitableAssetTypes      []string `json:"ClientSuitableAssetTypes"`
	HasAcceptedTermsAndConditions bool     `json:"HasAcceptedTermsAndConditions"`
	HasAccess                     bool     `json:"HasAccess"`
	IsDisabled                    bool     `json:"IsDisabled"`
	IsReady                       bool     `json:"IsReady"`
	IsUSExchangesRestricted       bool     `json:"IsUSExchangesRestricted"`
	SuitabilityLevel              string   `json:"SuitabilityLevel"`
}
