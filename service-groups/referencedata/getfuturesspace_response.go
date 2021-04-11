package referencedata

type GetFuturesSpaceResponse struct {
	BaseIdentifier string `json:"BaseIdentifier"`
	Elements       []struct {
		DaysToExpiry float64 `json:"DaysToExpiry"`
		ExpiryDate   string  `json:"ExpiryDate"`
		Symbol       string  `json:"Symbol"`
		Uic          float64 `json:"Uic"`
	} `json:"Elements"`
}
