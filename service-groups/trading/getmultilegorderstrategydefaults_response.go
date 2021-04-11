package trading

type GetMultiLegOrderStrategyDefaultsResponse struct {
	Legs []struct {
		Amount     float64 `json:"Amount"`
		AssetType  string  `json:"AssetType"`
		BuySell    string  `json:"BuySell"`
		OptionData struct {
			ExpiryDate  string  `json:"ExpiryDate"`
			PutCall     string  `json:"PutCall"`
			StrikePrice float64 `json:"StrikePrice"`
		} `json:"OptionData"`
		Uic float64 `json:"Uic"`
	} `json:"Legs"`
	StrategyType string `json:"StrategyType"`
}
