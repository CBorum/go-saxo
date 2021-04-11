package clientservices

type GetOrderStatesAsyncResponse struct {
	Data []struct {
		AccountID        string  `json:"AccountId"`
		ActivityTime     string  `json:"ActivityTime"`
		Amount           float64 `json:"Amount"`
		AssetType        string  `json:"AssetType"`
		BuySell          string  `json:"BuySell"`
		ClientID         string  `json:"ClientId"`
		CorrelationKey   string  `json:"CorrelationKey"`
		DisplayAndFormat struct {
			BarrierFormat string  `json:"BarrierFormat"`
			Currency      string  `json:"Currency"`
			Decimals      float64 `json:"Decimals"`
			Description   string  `json:"Description"`
			Format        string  `json:"Format"`
			OrderDecimals float64 `json:"OrderDecimals"`
			StrikeFormat  string  `json:"StrikeFormat"`
			Symbol        string  `json:"Symbol"`
		} `json:"DisplayAndFormat"`
		Duration struct {
			DurationType string `json:"DurationType"`
		} `json:"Duration"`
		HandledBy     string  `json:"HandledBy"`
		LogID         string  `json:"LogId"`
		OrderID       string  `json:"OrderId"`
		OrderRelation string  `json:"OrderRelation"`
		OrderType     string  `json:"OrderType"`
		Price         float64 `json:"Price"`
		Status        string  `json:"Status"`
		SubStatus     string  `json:"SubStatus"`
		Uic           float64 `json:"Uic"`
	} `json:"Data"`
	Next string `json:"__next"`
}
