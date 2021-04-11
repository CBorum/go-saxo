package portfolio

type GetInstrumentsExposuresResponse struct {
	Data []struct {
		Amount                 float64 `json:"Amount"`
		AssetType              string  `json:"AssetType"`
		AverageOpenPrice       float64 `json:"AverageOpenPrice"`
		CalculationReliability string  `json:"CalculationReliability"`
		CanBeClosed            bool    `json:"CanBeClosed"`
		DisplayAndFormat       struct {
			Currency    string  `json:"Currency"`
			Decimals    float64 `json:"Decimals"`
			Description string  `json:"Description"`
			Format      string  `json:"Format"`
			Symbol      string  `json:"Symbol"`
		} `json:"DisplayAndFormat"`
		NetPositionID string  `json:"NetPositionId"`
		Uic           float64 `json:"Uic"`
	} `json:"Data"`
	Next string `json:"__next"`
}
