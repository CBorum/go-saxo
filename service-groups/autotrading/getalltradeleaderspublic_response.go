package autotrading

type GetAllTradeLeadersPublicResponse struct {
	Data []struct {
		AccountPerformance struct {
			Allocation struct {
				TradesPerAssetType struct {
					ClosedTradesAllocations []struct {
						AssetClassType          string  `json:"AssetClassType"`
						CalculationBasis        float64 `json:"CalculationBasis"`
						CalculationDataType     float64 `json:"CalculationDataType"`
						PortfolioAllocationType float64 `json:"PortfolioAllocationType"`
						ReturnAttribution       float64 `json:"ReturnAttribution"`
						TradeCount              float64 `json:"TradeCount"`
						TradePct                float64 `json:"TradePct"`
						TradePercent            float64 `json:"TradePercent"`
					} `json:"ClosedTradesAllocations"`
				} `json:"TradesPerAssetType"`
				TradesPerInstrument struct {
					ClosedTradesAllocations []struct {
						AssetType               string  `json:"AssetType"`
						CalculationBasis        float64 `json:"CalculationBasis"`
						CalculationDataType     float64 `json:"CalculationDataType"`
						InstrumentDescription   string  `json:"InstrumentDescription"`
						InstrumentSymbol        string  `json:"InstrumentSymbol"`
						InstrumentUic           float64 `json:"InstrumentUic"`
						PortfolioAllocationType float64 `json:"PortfolioAllocationType"`
						ReturnAttribution       float64 `json:"ReturnAttribution"`
						TradeCount              float64 `json:"TradeCount"`
						TradePct                float64 `json:"TradePct"`
						TradePercent            float64 `json:"TradePercent"`
					} `json:"ClosedTradesAllocations"`
				} `json:"TradesPerInstrument"`
			} `json:"Allocation"`
			From                    string `json:"From"`
			InceptionDay            string `json:"InceptionDay"`
			LastTradeDay            string `json:"LastTradeDay"`
			Thru                    string `json:"Thru"`
			TimeWeightedPerformance struct {
				PerformanceFraction   float64 `json:"PerformanceFraction"`
				PerformanceKeyFigures struct {
					ClosedTradesCount        float64 `json:"ClosedTradesCount"`
					MaxDrawDownFraction      float64 `json:"MaxDrawDownFraction"`
					SampledStandardDeviation float64 `json:"SampledStandardDeviation"`
				} `json:"PerformanceKeyFigures"`
			} `json:"TimeWeightedPerformance"`
			TotalCashBalance                         float64       `json:"TotalCashBalance"`
			TotalOpenPositionsValue                  float64       `json:"TotalOpenPositionsValue"`
			TotalPositionsValuePerProductPerSecurity []interface{} `json:"TotalPositionsValuePerProductPerSecurity"`
		} `json:"AccountPerformance"`
		AllowMultipleInvestment     bool     `json:"AllowMultipleInvestment"`
		AllowPartialWithdrawal      bool     `json:"AllowPartialWithdrawal"`
		AutoFundTransferEnabled     bool     `json:"AutoFundTransferEnabled"`
		Category                    string   `json:"Category"`
		Currency                    string   `json:"Currency"`
		Description                 string   `json:"Description"`
		ExpectedCostExplanation     string   `json:"ExpectedCostExplanation"`
		FollowerAmountRemaining     float64  `json:"FollowerAmountRemaining"`
		IsInvestmentShieldSupported bool     `json:"IsInvestmentShieldSupported"`
		IsOpenForFollowers          bool     `json:"IsOpenForFollowers"`
		IsPortfolioMandateRequired  bool     `json:"IsPortfolioMandateRequired"`
		LeaderPartnerOptions        string   `json:"LeaderPartnerOptions"`
		LogoURL                     string   `json:"LogoUrl"`
		ManagementFee               float64  `json:"ManagementFee"`
		MinimumFundingAmount        float64  `json:"MinimumFundingAmount"`
		PerformanceFee              float64  `json:"PerformanceFee"`
		RiskLevel                   string   `json:"RiskLevel"`
		ShortDescription            string   `json:"ShortDescription"`
		StrategyLegalAssetTypes     []string `json:"StrategyLegalAssetTypes"`
		StrategyName                string   `json:"StrategyName"`
		SuitabilityLevel            string   `json:"SuitabilityLevel"`
		TotalExpectedCost           string   `json:"TotalExpectedCost"`
		TradeLeaderID               string   `json:"TradeLeaderId"`
		TradeLeaderOptions          string   `json:"TradeLeaderOptions"`
	} `json:"Data"`
}
