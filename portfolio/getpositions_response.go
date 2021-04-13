package portfolio

type GetPositionsResponse struct {
	Data []struct {
		NetPositionID string `json:"NetPositionId"`
		PositionBase  struct {
			AccountID                  string  `json:"AccountId"`
			Amount                     float64 `json:"Amount"`
			AssetType                  string  `json:"AssetType"`
			CanBeClosed                bool    `json:"CanBeClosed"`
			ClientID                   string  `json:"ClientId"`
			CloseConversionRateSettled bool    `json:"CloseConversionRateSettled"`
			ExecutionTimeOpen          string  `json:"ExecutionTimeOpen"`
			IsForceOpen                bool    `json:"IsForceOpen"`
			IsMarketOpen               bool    `json:"IsMarketOpen"`
			OpenPrice                  float64 `json:"OpenPrice"`
			SpotDate                   string  `json:"SpotDate"`
			Status                     string  `json:"Status"`
			Uic                        float64 `json:"Uic"`
			ValueDate                  string  `json:"ValueDate"`
		} `json:"PositionBase"`
		PositionID   string `json:"PositionId"`
		PositionView struct {
			Ask                             float64 `json:"Ask"`
			Bid                             float64 `json:"Bid"`
			CalculationReliability          string  `json:"CalculationReliability"`
			CurrentPrice                    float64 `json:"CurrentPrice"`
			CurrentPriceDelayMinutes        float64 `json:"CurrentPriceDelayMinutes"`
			CurrentPriceType                string  `json:"CurrentPriceType"`
			Exposure                        float64 `json:"Exposure"`
			ExposureCurrency                string  `json:"ExposureCurrency"`
			ExposureInBaseCurrency          float64 `json:"ExposureInBaseCurrency"`
			InstrumentPriceDayPercentChange float64 `json:"InstrumentPriceDayPercentChange"`
			ProfitLossOnTrade               float64 `json:"ProfitLossOnTrade"`
			ProfitLossOnTradeInBaseCurrency float64 `json:"ProfitLossOnTradeInBaseCurrency"`
			SettlementInstruction           struct {
				ActualRolloverAmount            float64 `json:"ActualRolloverAmount"`
				ActualSettlementAmount          float64 `json:"ActualSettlementAmount"`
				Amount                          float64 `json:"Amount"`
				IsSettlementInstructionsAllowed bool    `json:"IsSettlementInstructionsAllowed"`
				Month                           float64 `json:"Month"`
				SettlementType                  string  `json:"SettlementType"`
				Year                            float64 `json:"Year"`
			} `json:"SettlementInstruction"`
			TradeCostsTotal               float64 `json:"TradeCostsTotal"`
			TradeCostsTotalInBaseCurrency float64 `json:"TradeCostsTotalInBaseCurrency"`
		} `json:"PositionView"`
	} `json:"Data"`
	Next string `json:"__next"`
}
