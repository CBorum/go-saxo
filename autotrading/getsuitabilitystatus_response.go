package autotrading

type GetSuitabilityStatusResponse struct {
	CanGiveSuitabilityTest   bool     `json:"CanGiveSuitabilityTest"`
	LastSuitabilityCheckDate string   `json:"LastSuitabilityCheckDate"`
	OverallSuitability       string   `json:"OverallSuitability"`
	ProductAreaSuitability   []string `json:"ProductAreaSuitability"`
}
