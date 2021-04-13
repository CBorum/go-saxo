package referencedata

type GetAllExchangesResponse struct {
	Data []struct {
		CountryCode    string  `json:"CountryCode"`
		Currency       string  `json:"Currency"`
		ExchangeID     string  `json:"ExchangeId"`
		Mic            string  `json:"Mic"`
		Name           string  `json:"Name"`
		TimeZone       float64 `json:"TimeZone"`
		TimeZoneOffset string  `json:"TimeZoneOffset"`
	} `json:"Data"`
	Next string `json:"__next"`
}
