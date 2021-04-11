package portfolio

type GetOpenOrdersResponse struct {
	Data []struct {
		AccountID                string  `json:"AccountId"`
		AccountKey               string  `json:"AccountKey"`
		Amount                   float64 `json:"Amount"`
		AssetType                string  `json:"AssetType"`
		BuySell                  string  `json:"BuySell"`
		CalculationReliability   string  `json:"CalculationReliability"`
		ClientKey                string  `json:"ClientKey"`
		CurrentPrice             float64 `json:"CurrentPrice"`
		CurrentPriceDelayMinutes float64 `json:"CurrentPriceDelayMinutes"`
		CurrentPriceType         string  `json:"CurrentPriceType"`
		DistanceToMarket         float64 `json:"DistanceToMarket"`
		Duration                 struct {
			DurationType string `json:"DurationType"`
		} `json:"Duration"`
		IsForceOpen        bool    `json:"IsForceOpen"`
		IsMarketOpen       bool    `json:"IsMarketOpen"`
		MarketPrice        float64 `json:"MarketPrice"`
		NonTradableReason  string  `json:"NonTradableReason"`
		OpenOrderType      string  `json:"OpenOrderType"`
		OrderAmountType    string  `json:"OrderAmountType"`
		OrderID            string  `json:"OrderId"`
		OrderRelation      string  `json:"OrderRelation"`
		OrderTime          string  `json:"OrderTime"`
		Price              float64 `json:"Price"`
		PriceWithoutSpread float64 `json:"PriceWithoutSpread"`
		Status             string  `json:"Status"`
		Uic                float64 `json:"Uic"`
	} `json:"Data"`
	Next string `json:"__next"`
}
