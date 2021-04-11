package clientservices

type GetCasesResponse struct {
	Data []struct {
		CaseID                 string `json:"CaseId"`
		CaseLatestActivityTime string `json:"CaseLatestActivityTime"`
		CaseStatus             string `json:"CaseStatus"`
		ClientID               string `json:"ClientId"`
		CreatedOn              string `json:"CreatedOn"`
		CustomerName           string `json:"CustomerName"`
		FollowUpDateTime       string `json:"FollowUpDateTime"`
		HandledByPartner       bool   `json:"HandledByPartner"`
		IsEscalated            bool   `json:"IsEscalated"`
		Title                  string `json:"Title"`
	} `json:"Data"`
	Count float64 `json:"__count"`
}
