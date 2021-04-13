package trading

type AddMultiLegPricesSubscriptionAsyncResponse struct {
	ContextID         string  `json:"ContextId"`
	Format            string  `json:"Format"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Schema            string  `json:"Schema"`
	SchemaName        string  `json:"SchemaName"`
	Snapshot          struct {
		BuySell     string `json:"BuySell"`
		Commissions struct {
			CostBuy  float64 `json:"CostBuy"`
			CostSell float64 `json:"CostSell"`
		} `json:"Commissions"`
		Greeks struct {
			Delta  float64 `json:"Delta"`
			Gamma  float64 `json:"Gamma"`
			MidVol float64 `json:"MidVol"`
			Phi    float64 `json:"Phi"`
			Rho    float64 `json:"Rho"`
			Theta  float64 `json:"Theta"`
			Vega   float64 `json:"Vega"`
		} `json:"Greeks"`
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
		Legs        []struct {
			Amount    float64 `json:"Amount"`
			AssetType string  `json:"AssetType"`
			BuySell   string  `json:"BuySell"`
			Greeks    struct {
				Delta  float64 `json:"Delta"`
				Gamma  float64 `json:"Gamma"`
				MidVol float64 `json:"MidVol"`
				Phi    float64 `json:"Phi"`
				Rho    float64 `json:"Rho"`
				Theta  float64 `json:"Theta"`
				Vega   float64 `json:"Vega"`
			} `json:"Greeks"`
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
			LegID string `json:"LegId"`
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
			ToOpenClose string  `json:"ToOpenClose"`
			Uic         float64 `json:"Uic"`
		} `json:"Legs"`
		MarginImpactBuySell struct {
			Currency                      string  `json:"Currency"`
			InitialMarginAvailableBuy     float64 `json:"InitialMarginAvailableBuy"`
			InitialMarginAvailableCurrent float64 `json:"InitialMarginAvailableCurrent"`
			InitialMarginBuy              float64 `json:"InitialMarginBuy"`
			MaintenanceMarginBuy          float64 `json:"MaintenanceMarginBuy"`
		} `json:"MarginImpactBuySell"`
		PriceSource string `json:"PriceSource"`
		Quote       struct {
			Amount           float64 `json:"Amount"`
			Ask              float64 `json:"Ask"`
			Bid              float64 `json:"Bid"`
			DelayedByMinutes float64 `json:"DelayedByMinutes"`
			ErrorCode        string  `json:"ErrorCode"`
			Mid              float64 `json:"Mid"`
			PriceTypeAsk     string  `json:"PriceTypeAsk"`
			PriceTypeBid     string  `json:"PriceTypeBid"`
		} `json:"Quote"`
		StrategyType string `json:"StrategyType"`
	} `json:"Snapshot"`
	State string `json:"State"`
	Tag   string `json:"Tag"`
}
