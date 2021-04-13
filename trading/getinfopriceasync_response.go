package trading

type GetInfoPriceAsyncResponse struct {
	AssetType        string `json:"AssetType"`
	DisplayAndFormat struct {
		Currency      string  `json:"Currency"`
		Decimals      float64 `json:"Decimals"`
		Description   string  `json:"Description"`
		Format        string  `json:"Format"`
		OrderDecimals float64 `json:"OrderDecimals"`
		Symbol        string  `json:"Symbol"`
	} `json:"DisplayAndFormat"`
	HistoricalChanges struct {
		PercentChange1Month  float64 `json:"PercentChange1Month"`
		PercentChange2Months float64 `json:"PercentChange2Months"`
		PercentChange3Months float64 `json:"PercentChange3Months"`
		PercentChange6Months float64 `json:"PercentChange6Months"`
		PercentChangeDaily   float64 `json:"PercentChangeDaily"`
		PercentChangeWeekly  float64 `json:"PercentChangeWeekly"`
	} `json:"HistoricalChanges"`
	InstrumentPriceDetails struct {
		IsMarketOpen       bool   `json:"IsMarketOpen"`
		ShortTradeDisabled bool   `json:"ShortTradeDisabled"`
		ValueDate          string `json:"ValueDate"`
	} `json:"InstrumentPriceDetails"`
	LastUpdated string `json:"LastUpdated"`
	PriceInfo   struct {
		High          float64 `json:"High"`
		Low           float64 `json:"Low"`
		NetChange     float64 `json:"NetChange"`
		PercentChange float64 `json:"PercentChange"`
	} `json:"PriceInfo"`
	PriceInfoDetails struct {
		AskSize        float64 `json:"AskSize"`
		BidSize        float64 `json:"BidSize"`
		LastClose      float64 `json:"LastClose"`
		LastTraded     float64 `json:"LastTraded"`
		LastTradedSize float64 `json:"LastTradedSize"`
		Open           float64 `json:"Open"`
		Volume         float64 `json:"Volume"`
	} `json:"PriceInfoDetails"`
	Quote struct {
		Amount           float64 `json:"Amount"`
		Ask              float64 `json:"Ask"`
		Bid              float64 `json:"Bid"`
		DelayedByMinutes float64 `json:"DelayedByMinutes"`
		ErrorCode        string  `json:"ErrorCode"`
		Mid              float64 `json:"Mid"`
		PriceTypeAsk     string  `json:"PriceTypeAsk"`
		PriceTypeBid     string  `json:"PriceTypeBid"`
	} `json:"Quote"`
	Uic float64 `json:"Uic"`
}
