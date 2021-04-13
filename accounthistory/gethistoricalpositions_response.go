package accounthistory

type GetHistoricalPositionsResponse struct {
	Data []struct {
		AccountID            string `json:"AccountId"`
		AccountValueEndOfDay struct {
			AccountBalance    float64 `json:"AccountBalance"`
			CashTransfers     float64 `json:"CashTransfers"`
			Date              string  `json:"Date"`
			PositionsValue    float64 `json:"PositionsValue"`
			SecurityTransfers float64 `json:"SecurityTransfers"`
			TotalValue        float64 `json:"TotalValue"`
		} `json:"AccountValueEndOfDay"`
		Amount                             float64 `json:"Amount"`
		AmountAccountValueCloseRatio       string  `json:"AmountAccountValueCloseRatio"`
		AmountAccountValueOpenRatio        string  `json:"AmountAccountValueOpenRatio"`
		ClosingAssetType                   string  `json:"ClosingAssetType"`
		ClosingTradeDate                   string  `json:"ClosingTradeDate"`
		ClosingValueDate                   string  `json:"ClosingValueDate"`
		CopiedFrom                         string  `json:"CopiedFrom"`
		CorrelationType                    string  `json:"CorrelationType"`
		Decimals                           float64 `json:"Decimals"`
		ExecutionTimeClose                 string  `json:"ExecutionTimeClose"`
		ExecutionTimeOpen                  string  `json:"ExecutionTimeOpen"`
		FigureValue                        float64 `json:"FigureValue"`
		InstrumentCcyToAccountCcyRateClose float64 `json:"InstrumentCcyToAccountCcyRateClose"`
		InstrumentCcyToAccountCcyRateOpen  float64 `json:"InstrumentCcyToAccountCcyRateOpen"`
		InstrumentSymbol                   string  `json:"InstrumentSymbol"`
		LongShort                          struct {
			PresentationValue string `json:"PresentationValue"`
			Value             string `json:"Value"`
		} `json:"LongShort"`
		OpeningAssetType               string  `json:"OpeningAssetType"`
		OpeningTradeDate               string  `json:"OpeningTradeDate"`
		OpeningValueDate               string  `json:"OpeningValueDate"`
		PriceClose                     float64 `json:"PriceClose"`
		PriceGain                      float64 `json:"PriceGain"`
		PriceOpen                      float64 `json:"PriceOpen"`
		PricePct                       float64 `json:"PricePct"`
		ProfitLoss                     float64 `json:"ProfitLoss"`
		ProfitLossAccountValueFraction float64 `json:"ProfitLossAccountValueFraction"`
		Uic                            string  `json:"Uic"`
		ValueInAccountCurrencyClose    float64 `json:"ValueInAccountCurrencyClose"`
		ValueInAccountCurrencyOpen     float64 `json:"ValueInAccountCurrencyOpen"`
	} `json:"Data"`
}
