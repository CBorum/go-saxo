package trading

type RequestMarginImpactOnSubscriptionResponse struct {
	Data struct {
		AssetType        string `json:"AssetType"`
		DisplayAndFormat struct {
			Currency       string  `json:"Currency"`
			Decimals       float64 `json:"Decimals"`
			Description    string  `json:"Description"`
			Format         string  `json:"Format"`
			StrikeDecimals float64 `json:"StrikeDecimals"`
			StrikeFormat   string  `json:"StrikeFormat"`
			Symbol         string  `json:"Symbol"`
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
			AverageVolume          float64 `json:"AverageVolume"`
			Barrier                float64 `json:"Barrier"`
			ExpiryDate             string  `json:"ExpiryDate"`
			IsMarketOpen           bool    `json:"IsMarketOpen"`
			LowerBarrier           float64 `json:"LowerBarrier"`
			MidForwardPrice        float64 `json:"MidForwardPrice"`
			ShortTradeDisabled     bool    `json:"ShortTradeDisabled"`
			SpotAsk                float64 `json:"SpotAsk"`
			SpotBid                float64 `json:"SpotBid"`
			SpotDate               string  `json:"SpotDate"`
			SpreadStrikePriceLower float64 `json:"SpreadStrikePriceLower"`
			SpreadStrikePriceUpper float64 `json:"SpreadStrikePriceUpper"`
			StrikePrice            float64 `json:"StrikePrice"`
			UpperBarrier           float64 `json:"UpperBarrier"`
			ValueDate              string  `json:"ValueDate"`
		} `json:"InstrumentPriceDetails"`
		LastUpdated string `json:"LastUpdated"`
		Quote       struct {
			Amount           float64 `json:"Amount"`
			Ask              float64 `json:"Ask"`
			Bid              float64 `json:"Bid"`
			DelayedByMinutes float64 `json:"DelayedByMinutes"`
			ErrorCode        string  `json:"ErrorCode"`
			Mid              float64 `json:"Mid"`
			PriceTypeAsk     string  `json:"PriceTypeAsk"`
			PriceTypeBid     string  `json:"PriceTypeBid"`
			RFQState         string  `json:"RFQState"`
		} `json:"Quote"`
		Uic float64 `json:"Uic"`
	} `json:"Data"`
	PartitionNumber float64 `json:"PartitionNumber"`
	ReferenceID     string  `json:"ReferenceId"`
	Timestamp       string  `json:"Timestamp"`
	TotalPartitions float64 `json:"TotalPartitions"`
}
