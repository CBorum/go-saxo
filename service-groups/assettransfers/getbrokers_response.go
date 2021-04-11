package assettransfers

type GetBrokersResponse struct {
	Data []struct {
		Contact     string `json:"Contact"`
		CountryCode string `json:"CountryCode"`
		Email       string `json:"Email"`
		Name        string `json:"Name"`
		Phone       string `json:"Phone"`
	} `json:"Data"`
}
