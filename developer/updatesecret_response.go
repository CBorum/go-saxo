package developer

type UpdateSecretResponse struct {
	Secret     string  `json:"Secret"`
	SecretID   float64 `json:"SecretId"`
	ValidFrom  string  `json:"ValidFrom"`
	ValidUntil string  `json:"ValidUntil"`
}
