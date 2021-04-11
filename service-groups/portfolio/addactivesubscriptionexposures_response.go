package portfolio

type AddActiveSubscriptionExposuresResponse struct {
	ContextID         string  `json:"ContextId"`
	Format            string  `json:"Format"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Snapshot          struct {
		Data []struct {
			Amount                 float64 `json:"Amount"`
			AssetType              string  `json:"AssetType"`
			AverageOpenPrice       float64 `json:"AverageOpenPrice"`
			CalculationReliability string  `json:"CalculationReliability"`
			CanBeClosed            bool    `json:"CanBeClosed"`
			DisplayAndFormat       struct {
				Currency    string  `json:"Currency"`
				Decimals    float64 `json:"Decimals"`
				Description string  `json:"Description"`
				Format      string  `json:"Format"`
				Symbol      string  `json:"Symbol"`
			} `json:"DisplayAndFormat"`
			NetPositionID string  `json:"NetPositionId"`
			Uic           float64 `json:"Uic"`
		} `json:"Data"`
	} `json:"Snapshot"`
	State string `json:"State"`
}
