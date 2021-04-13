package root

type GetFeaturesResponse []struct {
	Available bool   `json:"Available"`
	Feature   string `json:"Feature"`
}
