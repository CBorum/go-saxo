package root

type PostFeaturesResponse struct {
	ContextID         string  `json:"ContextId"`
	Format            string  `json:"Format"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Snapshot          []struct {
		Available bool   `json:"Available"`
		Feature   string `json:"Feature"`
	} `json:"Snapshot"`
	State string `json:"State"`
}
