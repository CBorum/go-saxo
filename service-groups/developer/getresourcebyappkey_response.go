package developer

type GetResourceByAppKeyResponse struct {
	AppKey   string `json:"AppKey"`
	Branding []struct {
		Description string  `json:"Description"`
		ID          float64 `json:"Id"`
		Name        string  `json:"Name"`
	} `json:"Branding"`
	CreatedBy struct {
		Name    string `json:"Name"`
		UserKey string `json:"UserKey"`
	} `json:"CreatedBy"`
	Description string `json:"Description"`
	Endpoints   struct {
		AuthorizationEndpoint string `json:"AuthorizationEndpoint"`
		TokenEndpoint         string `json:"TokenEndpoint"`
	} `json:"Endpoints"`
	Flow             string `json:"Flow"`
	IsActive         bool   `json:"IsActive"`
	IsTradingEnabled bool   `json:"IsTradingEnabled"`
	Name             string `json:"Name"`
	RedirectUris     []struct {
		Branding struct {
			Description string  `json:"Description"`
			ID          float64 `json:"Id"`
			Name        string  `json:"Name"`
		} `json:"Branding"`
		Description   string  `json:"Description"`
		RedirectURIID float64 `json:"RedirectUriId"`
		URI           string  `json:"Uri"`
	} `json:"RedirectUris"`
	Secrets []struct {
		Secret     string  `json:"Secret"`
		SecretID   float64 `json:"SecretId"`
		ValidFrom  string  `json:"ValidFrom"`
		ValidUntil string  `json:"ValidUntil"`
	} `json:"Secrets"`
	Status string `json:"Status"`
}
