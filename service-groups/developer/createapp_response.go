package developer

type CreateAppResponse struct {
	AppKey    string `json:"AppKey"`
	Brandings []struct {
		Description string  `json:"Description"`
		ID          float64 `json:"Id"`
		Name        string  `json:"Name"`
	} `json:"Brandings"`
	CreatedBy struct {
		Name    string `json:"Name"`
		UserKey string `json:"UserKey"`
	} `json:"CreatedBy"`
	Description string `json:"Description"`
	Endpoints   struct {
		AuthorizationEndpoint string `json:"AuthorizationEndpoint"`
		TokenEndpoint         string `json:"TokenEndpoint"`
	} `json:"Endpoints"`
	Flow                         string `json:"Flow"`
	IsActive                     bool   `json:"IsActive"`
	IsTradingEnabled             bool   `json:"IsTradingEnabled"`
	ManualOrderIndicationDefault string `json:"ManualOrderIndicationDefault"`
	ManualOrderIndicator         string `json:"ManualOrderIndicator"`
	Name                         string `json:"Name"`
	Status                       string `json:"Status"`
}
