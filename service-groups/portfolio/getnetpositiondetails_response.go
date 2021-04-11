package portfolio

type GetNetPositionDetailsResponse struct {
	NetPositionBase struct {
		AccountID              string  `json:"AccountId"`
		Amount                 float64 `json:"Amount"`
		AssetType              string  `json:"AssetType"`
		CanBeClosed            bool    `json:"CanBeClosed"`
		ClientID               string  `json:"ClientId"`
		HasForceOpenPositions  bool    `json:"HasForceOpenPositions"`
		IsMarketOpen           bool    `json:"IsMarketOpen"`
		NonTradableReason      string  `json:"NonTradableReason"`
		NumberOfRelatedOrders  float64 `json:"NumberOfRelatedOrders"`
		OpenOrdersCount        float64 `json:"OpenOrdersCount"`
		OpenTriggerOrdersCount float64 `json:"OpenTriggerOrdersCount"`
		OpeningDirection       string  `json:"OpeningDirection"`
		PositionsAccount       string  `json:"PositionsAccount"`
		SinglePositionStatus   string  `json:"SinglePositionStatus"`
		Uic                    float64 `json:"Uic"`
		ValueDate              string  `json:"ValueDate"`
	} `json:"NetPositionBase"`
	NetPositionDetails struct {
		MarketValue               float64 `json:"MarketValue"`
		MarketValueInBaseCurrency float64 `json:"MarketValueInBaseCurrency"`
		NonTradableReason         string  `json:"NonTradableReason"`
	} `json:"NetPositionDetails"`
	NetPositionID   string `json:"NetPositionId"`
	NetPositionView struct {
		AverageOpenPrice                float64 `json:"AverageOpenPrice"`
		AverageOpenPriceIncludingCosts  float64 `json:"AverageOpenPriceIncludingCosts"`
		CalculationReliability          string  `json:"CalculationReliability"`
		CurrentPrice                    float64 `json:"CurrentPrice"`
		CurrentPriceDelayMinutes        float64 `json:"CurrentPriceDelayMinutes"`
		CurrentPriceType                string  `json:"CurrentPriceType"`
		Exposure                        float64 `json:"Exposure"`
		ExposureInBaseCurrency          float64 `json:"ExposureInBaseCurrency"`
		InstrumentPriceDayPercentChange float64 `json:"InstrumentPriceDayPercentChange"`
		PositionCount                   float64 `json:"PositionCount"`
		PositionsNotClosedCount         float64 `json:"PositionsNotClosedCount"`
		ProfitLossOnTrade               float64 `json:"ProfitLossOnTrade"`
		SettlementInstruction           struct {
			ActualRolloverAmount            float64 `json:"ActualRolloverAmount"`
			ActualSettlementAmount          float64 `json:"ActualSettlementAmount"`
			Amount                          float64 `json:"Amount"`
			IsSettlementInstructionsAllowed bool    `json:"IsSettlementInstructionsAllowed"`
			Month                           float64 `json:"Month"`
			SettlementType                  string  `json:"SettlementType"`
			Year                            float64 `json:"Year"`
		} `json:"SettlementInstruction"`
		Status                        string  `json:"Status"`
		TradeCostsTotal               float64 `json:"TradeCostsTotal"`
		TradeCostsTotalInBaseCurrency float64 `json:"TradeCostsTotalInBaseCurrency"`
	} `json:"NetPositionView"`
}
