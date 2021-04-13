package referencedata

type GetDetailsForManyInstrumentsResponse struct {
	Data []struct {
		AmountDecimals      float64 `json:"AmountDecimals"`
		AssetType           string  `json:"AssetType"`
		CurrencyCode        string  `json:"CurrencyCode"`
		DefaultAmount       float64 `json:"DefaultAmount"`
		DefaultSlippage     float64 `json:"DefaultSlippage"`
		DefaultSlippageType string  `json:"DefaultSlippageType"`
		Description         string  `json:"Description"`
		Exchange            struct {
			CountryCode string `json:"CountryCode"`
			ExchangeID  string `json:"ExchangeId"`
			Name        string `json:"Name"`
		} `json:"Exchange"`
		Format struct {
			Decimals      float64 `json:"Decimals"`
			Format        string  `json:"Format"`
			OrderDecimals float64 `json:"OrderDecimals"`
		} `json:"Format"`
		FxForwardMaxForwardDate string  `json:"FxForwardMaxForwardDate"`
		FxForwardMinForwardDate string  `json:"FxForwardMinForwardDate"`
		GroupID                 float64 `json:"GroupId"`
		IncrementSize           float64 `json:"IncrementSize"`
		IsTradable              bool    `json:"IsTradable"`
		OrderDistances          struct {
			EntryDefaultDistance          float64 `json:"EntryDefaultDistance"`
			EntryDefaultDistanceType      string  `json:"EntryDefaultDistanceType"`
			StopLimitDefaultDistance      float64 `json:"StopLimitDefaultDistance"`
			StopLimitDefaultDistanceType  string  `json:"StopLimitDefaultDistanceType"`
			StopLossDefaultDistance       float64 `json:"StopLossDefaultDistance"`
			StopLossDefaultDistanceType   string  `json:"StopLossDefaultDistanceType"`
			StopLossDefaultEnabled        bool    `json:"StopLossDefaultEnabled"`
			StopLossDefaultOrderType      string  `json:"StopLossDefaultOrderType"`
			TakeProfitDefaultDistance     float64 `json:"TakeProfitDefaultDistance"`
			TakeProfitDefaultDistanceType string  `json:"TakeProfitDefaultDistanceType"`
			TakeProfitDefaultEnabled      bool    `json:"TakeProfitDefaultEnabled"`
		} `json:"OrderDistances"`
		StandardAmounts     []float64 `json:"StandardAmounts"`
		SupportedOrderTypes []string  `json:"SupportedOrderTypes"`
		Symbol              string    `json:"Symbol"`
		TickSize            float64   `json:"TickSize"`
		TradableAs          []string  `json:"TradableAs"`
		TradableOn          []string  `json:"TradableOn"`
		TradingSignals      string    `json:"TradingSignals"`
		Uic                 float64   `json:"Uic"`
	} `json:"Data"`
	Next string `json:"__next"`
}
