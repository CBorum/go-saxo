package clientmanagement

type GetSignUpStatusResponse struct {
	ClientID        string `json:"ClientId"`
	ClientKey       string `json:"ClientKey"`
	Message         string `json:"Message"`
	OnboardingState string `json:"OnboardingState"`
}
