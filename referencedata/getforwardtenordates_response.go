package referencedata

type GetForwardTenorDatesResponse struct {
	Data []struct {
		Date  string  `json:"Date"`
		Unit  string  `json:"Unit"`
		Value float64 `json:"Value"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
