package clientservices

type GetCaseResponse struct {
	Case struct {
		Activities []struct {
			ActivityPriority string `json:"ActivityPriority"`
			ActivityTime     string `json:"ActivityTime"`
			ActivityType     string `json:"ActivityType"`
			From             struct {
				EmailAddress string `json:"EmailAddress"`
				Name         string `json:"Name"`
			} `json:"From"`
			PhoneNumber string `json:"PhoneNumber"`
			Status      string `json:"Status"`
		} `json:"Activities"`
		CaseID                 string `json:"CaseId"`
		CaseLatestActivityTime string `json:"CaseLatestActivityTime"`
		CaseStatus             string `json:"CaseStatus"`
		CaseType               string `json:"CaseType"`
		ClientID               string `json:"ClientId"`
		ClientName             string `json:"ClientName"`
		ClientSegment          string `json:"ClientSegment"`
		ContactPersonID        string `json:"ContactPersonId"`
		ContactPersonName      string `json:"ContactPersonName"`
		CreatedOn              string `json:"CreatedOn"`
		Description            string `json:"Description"`
		FirstResponseSent      bool   `json:"FirstResponseSent"`
		FollowUpDateTime       string `json:"FollowUpDateTime"`
		HandledByPartner       bool   `json:"HandledByPartner"`
		InternalComments       []struct {
			Comment   string `json:"Comment"`
			CreatedOn string `json:"CreatedOn"`
		} `json:"InternalComments"`
		IsEscalated bool `json:"IsEscalated"`
		Notes       []struct {
			Attachment struct {
				FileName string `json:"FileName"`
				MimeType string `json:"MimeType"`
			} `json:"Attachment"`
			CreatedOn  string `json:"CreatedOn"`
			NoteOrigin string `json:"NoteOrigin"`
			NoteText   string `json:"NoteText"`
			NoteTitle  string `json:"NoteTitle"`
		} `json:"Notes"`
		Origin            string  `json:"Origin"`
		ReassignmentCount float64 `json:"ReassignmentCount"`
		ServiceLanguage   string  `json:"ServiceLanguage"`
		ShowInPortal      bool    `json:"ShowInPortal"`
		Subject           string  `json:"Subject"`
		Title             string  `json:"Title"`
	} `json:"Case"`
}
