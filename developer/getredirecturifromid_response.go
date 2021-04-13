package developer

type GetRedirectUriFromIdResponse struct {
	Branding struct {
		Description string  `json:"Description"`
		ID          float64 `json:"Id"`
		Name        string  `json:"Name"`
	} `json:"Branding"`
	Description   string  `json:"Description"`
	RedirectURIID float64 `json:"RedirectUriId"`
	URI           string  `json:"Uri"`
}
