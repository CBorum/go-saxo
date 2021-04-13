package clientservices

type GetForInstrumentResponse struct {
	AccountCurrency  string `json:"AccountCurrency"`
	AmountCurrency   string `json:"AmountCurrency"`
	AssetType        string `json:"AssetType"`
	CommissionLimits []struct {
		Currency      string  `json:"Currency"`
		MinCommission float64 `json:"MinCommission"`
		OrderAction   string  `json:"OrderAction"`
		RateOnAmount  float64 `json:"RateOnAmount"`
	} `json:"CommissionLimits"`
	CurrencyCutSpreadRate float64 `json:"CurrencyCutSpreadRate"`
	ExposureLimits        []struct {
		Currency   string  `json:"Currency"`
		Identifier string  `json:"Identifier"`
		Level      string  `json:"Level"`
		RuleType   string  `json:"RuleType"`
		Value      float64 `json:"Value"`
	} `json:"ExposureLimits"`
	FinanceInterestMarkUp float64 `json:"FinanceInterestMarkUp"`
	ForexPriceBands       []struct {
		AutoExecuteEnabled bool    `json:"AutoExecuteEnabled"`
		AutoQuoteEnabled   bool    `json:"AutoQuoteEnabled"`
		DisplayDecimals    float64 `json:"DisplayDecimals"`
		MaxBand            float64 `json:"MaxBand"`
		Spread             float64 `json:"Spread"`
		SpreadType         string  `json:"SpreadType"`
		UpperBandLimit     float64 `json:"UpperBandLimit"`
	} `json:"ForexPriceBands"`
	InstrumentCurrency    string `json:"InstrumentCurrency"`
	IsTradable            bool   `json:"IsTradable"`
	MarginTierRequirement struct {
		Entries []struct {
			ExtraWeekMarginRate float64 `json:"ExtraWeekMarginRate"`
			InitialMarginRate   float64 `json:"InitialMarginRate"`
			IntraWeekMarginRate float64 `json:"IntraWeekMarginRate"`
			TierLowerBound      float64 `json:"TierLowerBound"`
		} `json:"Entries"`
		TierCurrency string `json:"TierCurrency"`
	} `json:"MarginTierRequirement"`
	MaxOrderAutoPlaceAmount float64  `json:"MaxOrderAutoPlaceAmount"`
	MaxStreamingAmount      float64  `json:"MaxStreamingAmount"`
	MinOrderSize            float64  `json:"MinOrderSize"`
	MinOrderSizeCurrency    string   `json:"MinOrderSizeCurrency"`
	Rating                  float64  `json:"Rating"`
	Rules                   []string `json:"Rules"`
	SpreadAsLowAs           float64  `json:"SpreadAsLowAs"`
	SwapInterestMarkUp      float64  `json:"SwapInterestMarkUp"`
	TicketFeeThreshold      float64  `json:"TicketFeeThreshold"`
	Uic                     float64  `json:"Uic"`
}
