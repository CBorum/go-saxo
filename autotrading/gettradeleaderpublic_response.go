package autotrading

type GetTradeLeaderPublicResponse struct {
	AccountPerformance struct {
		AccountSummary struct {
			AverageTradeDurationInMinutes float64  `json:"AverageTradeDurationInMinutes"`
			AverageTradesPerWeek          float64  `json:"AverageTradesPerWeek"`
			NumberOfDaysTraded            float64  `json:"NumberOfDaysTraded"`
			NumberOfLongTrades            float64  `json:"NumberOfLongTrades"`
			NumberOfShortTrades           float64  `json:"NumberOfShortTrades"`
			TopTradedInstruments          []string `json:"TopTradedInstruments"`
			TotalReturnFraction           float64  `json:"TotalReturnFraction"`
			TradedInstruments             []string `json:"TradedInstruments"`
			TradesTotalCount              float64  `json:"TradesTotalCount"`
			TradesWonCount                float64  `json:"TradesWonCount"`
			WinFraction                   float64  `json:"WinFraction"`
		} `json:"AccountSummary"`
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
		BalancePerformance struct {
			AccountBalanceTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"AccountBalanceTimeSeries"`
			AccountValueTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"AccountValueTimeSeries"`
			MonthlyProfitLossTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"MonthlyProfitLossTimeSeries"`
			SecurityTransferTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"SecurityTransferTimeSeries"`
			YearlyProfitLossTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"YearlyProfitLossTimeSeries"`
		} `json:"BalancePerformance"`
		BenchMark struct {
			AssetType       string  `json:"AssetType"`
			Description     string  `json:"Description"`
			DisplayHintType string  `json:"DisplayHintType"`
			Uic             float64 `json:"Uic"`
		} `json:"BenchMark"`
		BenchmarkPerformance []struct {
			Date  string  `json:"Date"`
			Value float64 `json:"Value"`
		} `json:"BenchmarkPerformance"`
		From                    string `json:"From"`
		InceptionDay            string `json:"InceptionDay"`
		LastTradeDay            string `json:"LastTradeDay"`
		Thru                    string `json:"Thru"`
		TimeWeightedPerformance struct {
			AccumulatedTimeWeightedTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"AccumulatedTimeWeightedTimeSeries"`
			MonthlyReturnTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"MonthlyReturnTimeSeries"`
			PerformanceFraction   float64 `json:"PerformanceFraction"`
			PerformanceKeyFigures struct {
				ClosedTradesCount float64 `json:"ClosedTradesCount"`
				DrawdownReport    struct {
					Drawdowns []struct {
						DaysCount      float64 `json:"DaysCount"`
						DepthInPercent float64 `json:"DepthInPercent"`
						FromDate       string  `json:"FromDate"`
						ThruDate       string  `json:"ThruDate"`
					} `json:"Drawdowns"`
					MaxDaysInDrawdownFromTop10Drawdowns float64 `json:"MaxDaysInDrawdownFromTop10Drawdowns"`
				} `json:"DrawdownReport"`
				LosingDaysFraction       float64 `json:"LosingDaysFraction"`
				MaxDrawDownFraction      float64 `json:"MaxDrawDownFraction"`
				ReturnFraction           float64 `json:"ReturnFraction"`
				SampledStandardDeviation float64 `json:"SampledStandardDeviation"`
				SharpeRatio              float64 `json:"SharpeRatio"`
				SortinoRatio             float64 `json:"SortinoRatio"`
			} `json:"PerformanceKeyFigures"`
			YearlyReturnTimeSeries []struct {
				Date  string  `json:"Date"`
				Value float64 `json:"Value"`
			} `json:"YearlyReturnTimeSeries"`
		} `json:"TimeWeightedPerformance"`
		TotalCashBalance            float64 `json:"TotalCashBalance"`
		TotalCashBalancePerCurrency []struct {
			StringValue string  `json:"StringValue"`
			Value       float64 `json:"Value"`
		} `json:"TotalCashBalancePerCurrency"`
		TotalOpenPositionsValue        float64 `json:"TotalOpenPositionsValue"`
		TotalPositionsValuePerCurrency []struct {
			StringValue string  `json:"StringValue"`
			Value       float64 `json:"Value"`
		} `json:"TotalPositionsValuePerCurrency"`
		TotalPositionsValuePerProductPerSecurity []struct {
			Description string  `json:"Description"`
			ProductName string  `json:"ProductName"`
			Symbol      string  `json:"Symbol"`
			Value       float64 `json:"Value"`
		} `json:"TotalPositionsValuePerProductPerSecurity"`
	} `json:"AccountPerformance"`
	AllowMultipleInvestment     bool    `json:"AllowMultipleInvestment"`
	AllowPartialWithdrawal      bool    `json:"AllowPartialWithdrawal"`
	AutoFundTransferEnabled     bool    `json:"AutoFundTransferEnabled"`
	Category                    string  `json:"Category"`
	Currency                    string  `json:"Currency"`
	FollowerAmountRemaining     float64 `json:"FollowerAmountRemaining"`
	IsInvestmentShieldSupported bool    `json:"IsInvestmentShieldSupported"`
	IsOpenForFollowers          bool    `json:"IsOpenForFollowers"`
	IsPortfolioMandateRequired  bool    `json:"IsPortfolioMandateRequired"`
	LeaderPartnerOptions        string  `json:"LeaderPartnerOptions"`
	ManagementFee               float64 `json:"ManagementFee"`
	MinimumFundingAmount        float64 `json:"MinimumFundingAmount"`
	PerformanceFee              float64 `json:"PerformanceFee"`
	Prospectus                  struct {
		Chcorp string `json:"CHCORP"`
		Chdvd  string `json:"CHDVD"`
		Chspi  string `json:"CHSPI"`
		CSBGC0 string `json:"CSBGC0"`
		CSBGC3 string `json:"CSBGC3"`
		CSBGC7 string `json:"CSBGC7"`
		Csgldc string `json:"CSGLDC"`
		Cspxj  string `json:"CSPXJ"`
		Cssmim string `json:"CSSMIM"`
		E20Y   string `json:"E20Y"`
		Eimi   string `json:"EIMI"`
		Embc   string `json:"EMBC"`
		Embe   string `json:"EMBE"`
		Emim   string `json:"EMIM"`
		Emuc   string `json:"EMUC"`
		Exxy   string `json:"EXXY"`
		Ghyc   string `json:"GHYC"`
		Ibcx   string `json:"IBCX"`
		Iega   string `json:"IEGA"`
		Iege   string `json:"IEGE"`
		Iemb   string `json:"IEMB"`
		Ijpc   string `json:"IJPC"`
		Imeu   string `json:"IMEU"`
		Infr   string `json:"INFR"`
		Ircp   string `json:"IRCP"`
		Iusc   string `json:"IUSC"`
		Iuse   string `json:"IUSE"`
		Iwdp   string `json:"IWDP"`
		Mveu   string `json:"MVEU"`
		Shyu   string `json:"SHYU"`
		Sjpa   string `json:"SJPA"`
	} `json:"Prospectus"`
	RiskLevel               string   `json:"RiskLevel"`
	StrategyLegalAssetTypes []string `json:"StrategyLegalAssetTypes"`
	StrategyName            string   `json:"StrategyName"`
	SuitabilityLevel        string   `json:"SuitabilityLevel"`
	TotalExpectedCost       string   `json:"TotalExpectedCost"`
	TradeLeaderID           string   `json:"TradeLeaderId"`
	TradeLeaderOptions      string   `json:"TradeLeaderOptions"`
}
