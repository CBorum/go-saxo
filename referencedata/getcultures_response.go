package referencedata

type GetCulturesResponse struct {
	Data []struct {
		CultureCode string `json:"CultureCode"`
		Name        string `json:"Name"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
