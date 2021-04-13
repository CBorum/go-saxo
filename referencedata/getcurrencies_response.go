package referencedata

type GetCurrenciesResponse struct {
	Data []struct {
		CurrencyCode string  `json:"CurrencyCode"`
		Decimals     float64 `json:"Decimals"`
		Name         string  `json:"Name"`
	} `json:"Data"`
	Next string `json:"__next"`
}
