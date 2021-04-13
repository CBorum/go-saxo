package charts

type GetChartDataAsyncResponse struct {
	ChartInfo struct {
		DelayedByMinutes float64 `json:"DelayedByMinutes"`
		ExchangeID       string  `json:"ExchangeId"`
		FirstSampleTime  string  `json:"FirstSampleTime"`
		Horizon          float64 `json:"Horizon"`
	} `json:"ChartInfo"`
	Data []struct {
		CloseAsk float64 `json:"CloseAsk"`
		CloseBid float64 `json:"CloseBid"`
		HighAsk  float64 `json:"HighAsk"`
		HighBid  float64 `json:"HighBid"`
		LowAsk   float64 `json:"LowAsk"`
		LowBid   float64 `json:"LowBid"`
		OpenAsk  float64 `json:"OpenAsk"`
		OpenBid  float64 `json:"OpenBid"`
		Time     string  `json:"Time"`
	} `json:"Data"`
	DataVersion      float64 `json:"DataVersion"`
	DisplayAndFormat struct {
		Currency    string  `json:"Currency"`
		Decimals    float64 `json:"Decimals"`
		Description string  `json:"Description"`
		Format      string  `json:"Format"`
		Symbol      string  `json:"Symbol"`
	} `json:"DisplayAndFormat"`
}
