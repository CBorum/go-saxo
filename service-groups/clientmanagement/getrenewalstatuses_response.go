package clientmanagement

type GetRenewalStatusesResponse struct {
	Data []struct {
		ClientID      string `json:"ClientId"`
		ClientKey     string `json:"ClientKey"`
		RenewalBy     string `json:"RenewalBy"`
		RenewalStatus string `json:"RenewalStatus"`
	} `json:"Data"`
	Count float64 `json:"__count"`
}
