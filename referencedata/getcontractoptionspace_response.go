package referencedata

type GetContractOptionSpaceResponse struct {
	AmountDecimals                float64 `json:"AmountDecimals"`
	AssetType                     string  `json:"AssetType"`
	CanParticipateInMultiLegOrder bool    `json:"CanParticipateInMultiLegOrder"`
	ContractSize                  float64 `json:"ContractSize"`
	CurrencyCode                  string  `json:"CurrencyCode"`
	DefaultAmount                 float64 `json:"DefaultAmount"`
	Description                   string  `json:"Description"`
	Exchange                      struct {
		CountryCode     string `json:"CountryCode"`
		ExchangeID      string `json:"ExchangeId"`
		Name            string `json:"Name"`
		PriceSourceName string `json:"PriceSourceName"`
	} `json:"Exchange"`
	ExerciseStyle string `json:"ExerciseStyle"`
	Format        struct {
		Decimals       float64 `json:"Decimals"`
		OrderDecimals  float64 `json:"OrderDecimals"`
		StrikeDecimals float64 `json:"StrikeDecimals"`
	} `json:"Format"`
	GroupID       float64 `json:"GroupId"`
	IncrementSize float64 `json:"IncrementSize"`
	IsTradable    bool    `json:"IsTradable"`
	LotSize       float64 `json:"LotSize"`
	LotSizeType   string  `json:"LotSizeType"`
	OptionRootID  float64 `json:"OptionRootId"`
	OptionSpace   []struct {
		DisplayDaysToExpiry        float64 `json:"DisplayDaysToExpiry"`
		DisplayDaysToLastTradeDate float64 `json:"DisplayDaysToLastTradeDate"`
		DisplayExpiry              string  `json:"DisplayExpiry"`
		Expiry                     string  `json:"Expiry"`
		LastTradeDate              string  `json:"LastTradeDate"`
		SpecificOptions            []struct {
			PutCall       string  `json:"PutCall"`
			StrikePrice   float64 `json:"StrikePrice"`
			TradingStatus string  `json:"TradingStatus"`
			Uic           float64 `json:"Uic"`
			UnderlyingUic float64 `json:"UnderlyingUic"`
		} `json:"SpecificOptions"`
	} `json:"OptionSpace"`
	OrderDistances struct {
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
	PriceToContractFactor float64 `json:"PriceToContractFactor"`
	RelatedInstruments    []struct {
		AssetType string  `json:"AssetType"`
		Uic       float64 `json:"Uic"`
	} `json:"RelatedInstruments"`
	RelatedOptionRootsEnhanced []struct {
		AssetType    string  `json:"AssetType"`
		OptionRootID float64 `json:"OptionRootId"`
	} `json:"RelatedOptionRootsEnhanced"`
	SettlementStyle     string    `json:"SettlementStyle"`
	StandardAmounts     []float64 `json:"StandardAmounts"`
	SupportedOrderTypes []string  `json:"SupportedOrderTypes"`
	Symbol              string    `json:"Symbol"`
	TickSize            float64   `json:"TickSize"`
	TradableOn          []string  `json:"TradableOn"`
	UnderlyingAssetType string    `json:"UnderlyingAssetType"`
}
