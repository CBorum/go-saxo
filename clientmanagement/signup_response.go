package clientmanagement

type SignUpResponse struct {
	ClientID  string `json:"ClientId"`
	ClientKey string `json:"ClientKey"`
	SignUpID  string `json:"SignUpId"`
}
