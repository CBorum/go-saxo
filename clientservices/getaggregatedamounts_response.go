package clientservices

type GetAggregatedAmountsResponse struct {
	Data []struct {
		AccountCurrency                 string  `json:"AccountCurrency"`
		AffectsBalance                  bool    `json:"AffectsBalance"`
		Amount                          float64 `json:"Amount"`
		AmountAccountCurrency           float64 `json:"AmountAccountCurrency"`
		AmountClass                     string  `json:"AmountClass"`
		AmountClientCurrency            float64 `json:"AmountClientCurrency"`
		AmountSubClass                  string  `json:"AmountSubClass"`
		AmountTypeID                    float64 `json:"AmountTypeId"`
		AmountTypeName                  string  `json:"AmountTypeName"`
		AmountUSD                       float64 `json:"AmountUSD"`
		AssetType                       string  `json:"AssetType"`
		BookingAccountID                string  `json:"BookingAccountId"`
		ClientCurrency                  string  `json:"ClientCurrency"`
		CostClass                       string  `json:"CostClass"`
		CostSubClass                    string  `json:"CostSubClass"`
		Date                            string  `json:"Date"`
		InstrumentDescription           string  `json:"InstrumentDescription"`
		InstrumentSectorName            string  `json:"InstrumentSectorName"`
		InstrumentSectorTypeID          float64 `json:"InstrumentSectorTypeId"`
		InstrumentSubType               string  `json:"InstrumentSubType"`
		InstrumentSymbol                string  `json:"InstrumentSymbol"`
		RootInstrumentSectorName        string  `json:"RootInstrumentSectorName"`
		RootInstrumentSectorTypeID      float64 `json:"RootInstrumentSectorTypeId"`
		Uic                             float64 `json:"Uic"`
		UnderlyingInstrumentAssetType   string  `json:"UnderlyingInstrumentAssetType"`
		UnderlyingInstrumentDescription string  `json:"UnderlyingInstrumentDescription"`
		UnderlyingInstrumentSubType     string  `json:"UnderlyingInstrumentSubType"`
		UnderlyingInstrumentSymbol      string  `json:"UnderlyingInstrumentSymbol"`
		UnderlyingInstrumentUic         float64 `json:"UnderlyingInstrumentUic"`
	} `json:"Data"`
	Next string `json:"__next"`
}
