package portfolio

type GetClientMeResponse struct {
	ClientID                            string   `json:"ClientId"`
	ClientKey                           string   `json:"ClientKey"`
	CurrencyDecimals                    float64  `json:"CurrencyDecimals"`
	DefaultAccountID                    string   `json:"DefaultAccountId"`
	DefaultAccountKey                   string   `json:"DefaultAccountKey"`
	DefaultCurrency                     string   `json:"DefaultCurrency"`
	ForceOpenDefaultValue               bool     `json:"ForceOpenDefaultValue"`
	IsMarginTradingAllowed              bool     `json:"IsMarginTradingAllowed"`
	IsVariationMarginEligible           bool     `json:"IsVariationMarginEligible"`
	LegalAssetTypes                     []string `json:"LegalAssetTypes"`
	LegalAssetTypesAreIndicative        bool     `json:"LegalAssetTypesAreIndicative"`
	MarginCalculationMethod             string   `json:"MarginCalculationMethod"`
	Name                                string   `json:"Name"`
	PositionNettingMethod               string   `json:"PositionNettingMethod"`
	PositionNettingMode                 string   `json:"PositionNettingMode"`
	PositionNettingProfile              string   `json:"PositionNettingProfile"`
	ReduceExposureOnly                  bool     `json:"ReduceExposureOnly"`
	SupportsAccountValueProtectionLimit bool     `json:"SupportsAccountValueProtectionLimit"`
}
