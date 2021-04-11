package developer

type GetSecretsResponse struct {
	Data []struct {
		Secret     string  `json:"Secret"`
		SecretID   float64 `json:"SecretId"`
		ValidFrom  string  `json:"ValidFrom"`
		ValidUntil string  `json:"ValidUntil"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
