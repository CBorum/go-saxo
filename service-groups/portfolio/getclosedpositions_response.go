package portfolio

type GetClosedPositionsResponse struct {
	Data []struct {
		ClosedPosition struct {
			AccountID                                    string  `json:"AccountId"`
			Amount                                       float64 `json:"Amount"`
			AssetType                                    string  `json:"AssetType"`
			BuyOrSell                                    string  `json:"BuyOrSell"`
			ClientID                                     string  `json:"ClientId"`
			ClosedProfitLoss                             float64 `json:"ClosedProfitLoss"`
			ClosedProfitLossInBaseCurrency               float64 `json:"ClosedProfitLossInBaseCurrency"`
			ClosingMarketValue                           float64 `json:"ClosingMarketValue"`
			ClosingMarketValueInBaseCurrency             float64 `json:"ClosingMarketValueInBaseCurrency"`
			ClosingMethod                                string  `json:"ClosingMethod"`
			ClosingPositionID                            string  `json:"ClosingPositionId"`
			ClosingPrice                                 float64 `json:"ClosingPrice"`
			ConversionRateInstrumentToBaseSettledClosing bool    `json:"ConversionRateInstrumentToBaseSettledClosing"`
			ConversionRateInstrumentToBaseSettledOpening bool    `json:"ConversionRateInstrumentToBaseSettledOpening"`
			ExecutionTimeClose                           string  `json:"ExecutionTimeClose"`
			ExecutionTimeOpen                            string  `json:"ExecutionTimeOpen"`
			OpenPrice                                    float64 `json:"OpenPrice"`
			OpeningPositionID                            string  `json:"OpeningPositionId"`
			Uic                                          float64 `json:"Uic"`
		} `json:"ClosedPosition"`
		ClosedPositionUniqueID string `json:"ClosedPositionUniqueId"`
		NetPositionID          string `json:"NetPositionId"`
	} `json:"Data"`
	Next string `json:"__next"`
}
