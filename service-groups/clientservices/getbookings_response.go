package clientservices

type GetBookingsResponse struct {
	Data []struct {
		AccountCurrency                 string  `json:"AccountCurrency"`
		AccountID                       string  `json:"AccountId"`
		Amount                          float64 `json:"Amount"`
		AmountAccountCurrency           float64 `json:"AmountAccountCurrency"`
		AmountClass                     string  `json:"AmountClass"`
		AmountClientCurrency            float64 `json:"AmountClientCurrency"`
		AmountSubClass                  string  `json:"AmountSubClass"`
		AmountUSD                       float64 `json:"AmountUSD"`
		AssetType                       string  `json:"AssetType"`
		BkAmountID                      string  `json:"BkAmountId"`
		BkAmountType                    string  `json:"BkAmountType"`
		BkAmountTypeID                  string  `json:"BkAmountTypeId"`
		CaMasterRecordID                string  `json:"CaMasterRecordId"`
		ClientCurrency                  string  `json:"ClientCurrency"`
		ConversionRate                  float64 `json:"ConversionRate"`
		ConversionRateAccountCurrency   float64 `json:"ConversionRateAccountCurrency"`
		ConversionRateClientCurrency    float64 `json:"ConversionRateClientCurrency"`
		ConversionRateUSD               float64 `json:"ConversionRateUSD"`
		CostClass                       string  `json:"CostClass"`
		CostSubClass                    string  `json:"CostSubClass"`
		Currency                        string  `json:"Currency"`
		Date                            string  `json:"Date"`
		InstrumentDescription           string  `json:"InstrumentDescription"`
		InstrumentSectorName            string  `json:"InstrumentSectorName"`
		InstrumentSectorTypeID          float64 `json:"InstrumentSectorTypeId"`
		InstrumentSubType               string  `json:"InstrumentSubType"`
		InstrumentSymbol                string  `json:"InstrumentSymbol"`
		RelatedPositionID               string  `json:"RelatedPositionId"`
		RelatedTradeID                  string  `json:"RelatedTradeId"`
		RootInstrumentSectorName        string  `json:"RootInstrumentSectorName"`
		RootInstrumentSectorTypeID      float64 `json:"RootInstrumentSectorTypeId"`
		Uic                             float64 `json:"Uic"`
		UnderlyingInstrumentAssetType   string  `json:"UnderlyingInstrumentAssetType"`
		UnderlyingInstrumentDescription string  `json:"UnderlyingInstrumentDescription"`
		UnderlyingInstrumentSectorName  string  `json:"UnderlyingInstrumentSectorName"`
		UnderlyingInstrumentSubType     string  `json:"UnderlyingInstrumentSubType"`
		UnderlyingInstrumentSymbol      string  `json:"UnderlyingInstrumentSymbol"`
		UnderlyingInstrumentUic         float64 `json:"UnderlyingInstrumentUic"`
		ValueDate                       string  `json:"ValueDate"`
	} `json:"Data"`
	Next string `json:"__next"`
}
