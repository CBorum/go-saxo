package clientmanagement

type CompleteApplicationResponse struct {
	AccountDetails []struct {
		AccountType     string `json:"AccountType"`
		CurrencyIsoCode string `json:"CurrencyIsoCode"`
		Iban            string `json:"Iban"`
	} `json:"AccountDetails"`
	ApplicationStatus string `json:"ApplicationStatus"`
}
