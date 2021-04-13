package clientmanagement

type GetSignupOptionsResponse struct {
	Data []struct {
		PropertyName string `json:"PropertyName"`
		ValuePairs   []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"ValuePairs"`
	} `json:"Data"`
}
