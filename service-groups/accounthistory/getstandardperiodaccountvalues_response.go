package accounthistory

type GetStandardPeriodAccountValuesResponse struct {
	Data []struct {
		AccountValue      float64 `json:"AccountValue"`
		AccountValueMonth float64 `json:"AccountValueMonth"`
		AccountValueYear  float64 `json:"AccountValueYear"`
		Key               string  `json:"Key"`
		KeyType           string  `json:"KeyType"`
	} `json:"Data"`
}
