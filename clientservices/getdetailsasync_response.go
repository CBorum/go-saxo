package clientservices

type GetDetailsAsyncResponse struct {
	Data []struct {
		AccountCurrency                 string  `json:"AccountCurrency"`
		AccountCurrencyDecimals         float64 `json:"AccountCurrencyDecimals"`
		AccountID                       string  `json:"AccountId"`
		AdjustedTradeDate               string  `json:"AdjustedTradeDate"`
		Amount                          float64 `json:"Amount"`
		AssetType                       string  `json:"AssetType"`
		ClientCurrency                  string  `json:"ClientCurrency"`
		Direction                       string  `json:"Direction"`
		ExchangeDescription             string  `json:"ExchangeDescription"`
		InitialMargin                   float64 `json:"InitialMargin"`
		InstrumentCurrencyDecimal       float64 `json:"InstrumentCurrencyDecimal"`
		InstrumentDescription           string  `json:"InstrumentDescription"`
		InstrumentSectorName            string  `json:"InstrumentSectorName"`
		InstrumentSectorTypeID          float64 `json:"InstrumentSectorTypeId"`
		InstrumentSymbol                string  `json:"InstrumentSymbol"`
		MaintenanceMargin               float64 `json:"MaintenanceMargin"`
		OrderID                         string  `json:"OrderId"`
		Price                           float64 `json:"Price"`
		RootInstrumentSectorName        string  `json:"RootInstrumentSectorName"`
		RootInstrumentSectorTypeID      float64 `json:"RootInstrumentSectorTypeId"`
		Strike                          float64 `json:"Strike"`
		ToOpenOrClose                   string  `json:"ToOpenOrClose"`
		TradeBarrierEventStatus         bool    `json:"TradeBarrierEventStatus"`
		TradeDate                       string  `json:"TradeDate"`
		TradeEventType                  string  `json:"TradeEventType"`
		TradeExecutionTime              string  `json:"TradeExecutionTime"`
		TradeID                         string  `json:"TradeId"`
		TradeType                       string  `json:"TradeType"`
		TradedValue                     float64 `json:"TradedValue"`
		Uic                             float64 `json:"Uic"`
		UnderlyingInstrumentDescription string  `json:"UnderlyingInstrumentDescription"`
		UnderlyingInstrumentSymbol      string  `json:"UnderlyingInstrumentSymbol"`
		ValueDate                       string  `json:"ValueDate"`
		Venue                           string  `json:"Venue"`
	} `json:"Data"`
	Next string `json:"__next"`
}
