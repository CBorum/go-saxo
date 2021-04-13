package referencedata

type GetSummariesResponse struct {
	Data []struct {
		AssetType      string   `json:"AssetType"`
		Description    string   `json:"Description"`
		ExchangeID     string   `json:"ExchangeId"`
		GroupID        float64  `json:"GroupId"`
		Identifier     float64  `json:"Identifier"`
		IsKeywordMatch bool     `json:"IsKeywordMatch"`
		PrimaryListing float64  `json:"PrimaryListing"`
		SummaryType    string   `json:"SummaryType"`
		Symbol         string   `json:"Symbol"`
		TradableAs     []string `json:"TradableAs"`
	} `json:"Data"`
	Next string `json:"__next"`
}
