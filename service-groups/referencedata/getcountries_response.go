package referencedata

type GetCountriesResponse struct {
	Data []struct {
		A3          string  `json:"A3"`
		CountryCode string  `json:"CountryCode"`
		Name        string  `json:"Name"`
		Numeric     float64 `json:"Numeric"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
