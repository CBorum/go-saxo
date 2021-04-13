package valueadd

type GetAllAlertDefinitionsFilteredByStateResponse struct {
	Data []struct {
		AccountID         string  `json:"AccountId"`
		AlertDefinitionID string  `json:"AlertDefinitionId"`
		AssetType         string  `json:"AssetType"`
		Comment           string  `json:"Comment"`
		ExpiryDate        string  `json:"ExpiryDate"`
		IsRecurring       bool    `json:"IsRecurring"`
		Operator          string  `json:"Operator"`
		PriceVariable     string  `json:"PriceVariable"`
		State             string  `json:"State"`
		TargetValue       float64 `json:"TargetValue"`
		Uic               float64 `json:"Uic"`
		UserID            string  `json:"UserId"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
