package valueadd

type GetUserSettingsResponse struct {
	EmailAddress          string `json:"EmailAddress"`
	EmailAddressValidated bool   `json:"EmailAddressValidated"`
	NotifyWithMail        bool   `json:"NotifyWithMail"`
	NotifyWithPopup       bool   `json:"NotifyWithPopup"`
	Sound                 string `json:"Sound"`
}
