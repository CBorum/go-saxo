package clientservices

type GetForContractOptionResponse struct {
	AccountCurrency string `json:"AccountCurrency"`
	AssetType       string `json:"AssetType"`
	CarryingCost    struct {
		MarkUpRate float64 `json:"MarkUpRate"`
	} `json:"CarryingCost"`
	CommissionLimits []struct {
		Currency    string  `json:"Currency"`
		OrderAction string  `json:"OrderAction"`
		PerUnitRate float64 `json:"PerUnitRate"`
	} `json:"CommissionLimits"`
	CurrencyCutSpreadRate float64 `json:"CurrencyCutSpreadRate"`
	ExerciseCutOffTime    string  `json:"ExerciseCutOffTime"`
	ExerciseCutoffTime    string  `json:"ExerciseCutoffTime"`
	ExpirationTime        string  `json:"ExpirationTime"`
	ExpiryTime            string  `json:"ExpiryTime"`
	ExposureLimits        []struct {
		Identifier string  `json:"Identifier"`
		Level      string  `json:"Level"`
		RuleType   string  `json:"RuleType"`
		Value      float64 `json:"Value"`
	} `json:"ExposureLimits"`
	FinanceInterestMarkUp float64 `json:"FinanceInterestMarkUp"`
	HoldingFee            struct {
		FromPeriodInDays float64 `json:"FromPeriodInDays"`
		Value            float64 `json:"Value"`
	} `json:"HoldingFee"`
	InstrumentCurrency string `json:"InstrumentCurrency"`
	IsTradable         bool   `json:"IsTradable"`
	MarginRequirement  struct {
		MinimumUnderlyingValue           float64 `json:"MinimumUnderlyingValue"`
		Premium                          float64 `json:"Premium"`
		TradingProfile                   string  `json:"TradingProfile"`
		UnderlyingValueOvernightExposure float64 `json:"UnderlyingValueOvernightExposure"`
	} `json:"MarginRequirement"`
	Rating             float64 `json:"Rating"`
	SettlementStyle    string  `json:"SettlementStyle"`
	SwapInterestMarkUp float64 `json:"SwapInterestMarkUp"`
	Uic                float64 `json:"Uic"`
}
