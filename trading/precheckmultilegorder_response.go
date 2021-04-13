package trading

type PreCheckMultilegOrderResponse struct {
	Cost struct {
		Commission        float64 `json:"Commission"`
		ExchangeFee       float64 `json:"ExchangeFee"`
		GuaranteedStopFee float64 `json:"GuaranteedStopFee"`
		StampDuty         float64 `json:"StampDuty"`
	} `json:"Cost"`
	EstimatedCashRequired               float64 `json:"EstimatedCashRequired"`
	EstimatedCashRequiredCurrency       string  `json:"EstimatedCashRequiredCurrency"`
	EstimatedTotalCost                  float64 `json:"EstimatedTotalCost"`
	EstimatedTotalCostInAccountCurrency float64 `json:"EstimatedTotalCostInAccountCurrency"`
	InstrumentToAccountConversionRate   float64 `json:"InstrumentToAccountConversionRate"`
	MarginImpactBuySell                 struct {
		Currency                      string  `json:"Currency"`
		InitialMarginAvailableBuy     float64 `json:"InitialMarginAvailableBuy"`
		InitialMarginAvailableCurrent float64 `json:"InitialMarginAvailableCurrent"`
		InitialMarginBuy              float64 `json:"InitialMarginBuy"`
		MaintenanceMarginBuy          float64 `json:"MaintenanceMarginBuy"`
	} `json:"MarginImpactBuySell"`
	PreCheckResult string `json:"PreCheckResult"`
}
