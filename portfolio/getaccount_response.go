package portfolio

type GetAccountResponse struct {
	AccountGroupKey                       string   `json:"AccountGroupKey"`
	AccountID                             string   `json:"AccountId"`
	AccountKey                            string   `json:"AccountKey"`
	AccountType                           string   `json:"AccountType"`
	Active                                bool     `json:"Active"`
	CanUseCashPositionsAsMarginCollateral bool     `json:"CanUseCashPositionsAsMarginCollateral"`
	CfdBorrowingCostsActive               bool     `json:"CfdBorrowingCostsActive"`
	ClientID                              string   `json:"ClientId"`
	ClientKey                             string   `json:"ClientKey"`
	CreationDate                          string   `json:"CreationDate"`
	Currency                              string   `json:"Currency"`
	CurrencyDecimals                      float64  `json:"CurrencyDecimals"`
	DirectMarketAccess                    bool     `json:"DirectMarketAccess"`
	IndividualMargining                   bool     `json:"IndividualMargining"`
	IsCurrencyConversionAtSettlementTime  bool     `json:"IsCurrencyConversionAtSettlementTime"`
	IsMarginTradingAllowed                bool     `json:"IsMarginTradingAllowed"`
	IsShareable                           bool     `json:"IsShareable"`
	IsTrialAccount                        bool     `json:"IsTrialAccount"`
	LegalAssetTypes                       []string `json:"LegalAssetTypes"`
	MarginCalculationMethod               string   `json:"MarginCalculationMethod"`
	Sharing                               []string `json:"Sharing"`
	SupportsAccountValueProtectionLimit   bool     `json:"SupportsAccountValueProtectionLimit"`
	UseCashPositionsAsMarginCollateral    bool     `json:"UseCashPositionsAsMarginCollateral"`
}
