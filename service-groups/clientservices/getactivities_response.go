package clientservices

type GetActivitiesResponse struct {
	Data []struct {
		AccountID      string `json:"AccountId"`
		ClientID       string `json:"ClientId"`
		CorrelationKey string `json:"CorrelationKey"`
		Created        string `json:"Created"`
		Description    string `json:"Description"`
		LogID          string `json:"LogId"`
		PostingID      string `json:"PostingId"`
		UserID         string `json:"UserId"`
	} `json:"Data"`
	Next string `json:"__next"`
}
