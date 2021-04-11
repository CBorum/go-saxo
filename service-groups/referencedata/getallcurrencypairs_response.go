package referencedata

type GetAllCurrencyPairsResponse struct {
	Data []struct {
		CurrencyCode      string `json:"CurrencyCode"`
		RelatedCurrencies []struct {
			CurrencyCode string  `json:"CurrencyCode"`
			Uic          float64 `json:"Uic"`
		} `json:"RelatedCurrencies"`
	} `json:"Data"`
}
