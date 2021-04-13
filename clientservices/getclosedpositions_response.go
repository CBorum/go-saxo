package clientservices

type GetClosedPositionsResponse struct {
	Data []struct {
		AccountCurrency                        string  `json:"AccountCurrency"`
		AccountCurrencyDecimals                float64 `json:"AccountCurrencyDecimals"`
		AccountID                              string  `json:"AccountId"`
		Amount                                 float64 `json:"Amount"`
		AmountClose                            float64 `json:"AmountClose"`
		AmountOpen                             float64 `json:"AmountOpen"`
		AssetType                              string  `json:"AssetType"`
		ClosePositionID                        string  `json:"ClosePositionId"`
		ClosePrice                             float64 `json:"ClosePrice"`
		ExchangeDescription                    string  `json:"ExchangeDescription"`
		ISINCode                               string  `json:"ISINCode"`
		InstrumentDescription                  string  `json:"InstrumentDescription"`
		InstrumentSectorName                   string  `json:"InstrumentSectorName"`
		InstrumentSectorTypeID                 float64 `json:"InstrumentSectorTypeId"`
		InstrumentSymbol                       string  `json:"InstrumentSymbol"`
		OpenPositionID                         string  `json:"OpenPositionId"`
		OpenPrice                              float64 `json:"OpenPrice"`
		PnLAccountCurrency                     float64 `json:"PnLAccountCurrency"`
		PnLClientCurrency                      float64 `json:"PnLClientCurrency"`
		PnLUSD                                 float64 `json:"PnLUSD"`
		RootInstrumentSectorName               string  `json:"RootInstrumentSectorName"`
		RootInstrumentSectorTypeID             float64 `json:"RootInstrumentSectorTypeId"`
		TotalBookedOnClosingLegAccountCurrency float64 `json:"TotalBookedOnClosingLegAccountCurrency"`
		TotalBookedOnClosingLegClientCurrency  float64 `json:"TotalBookedOnClosingLegClientCurrency"`
		TotalBookedOnClosingLegUSD             float64 `json:"TotalBookedOnClosingLegUSD"`
		TotalBookedOnOpeningLegAccountCurrency float64 `json:"TotalBookedOnOpeningLegAccountCurrency"`
		TotalBookedOnOpeningLegClientCurrency  float64 `json:"TotalBookedOnOpeningLegClientCurrency"`
		TotalBookedOnOpeningLegUSD             float64 `json:"TotalBookedOnOpeningLegUSD"`
		TradeDate                              string  `json:"TradeDate"`
		TradeDateClose                         string  `json:"TradeDateClose"`
		TradeDateOpen                          string  `json:"TradeDateOpen"`
		UnderlyingInstrumentDescription        string  `json:"UnderlyingInstrumentDescription"`
		UnderlyingInstrumentSymbol             string  `json:"UnderlyingInstrumentSymbol"`
	} `json:"Data"`
	Next string `json:"__next"`
}
