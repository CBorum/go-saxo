package clientservices

type GetTradingConditionCostResponse struct {
	AccountCurrency string  `json:"AccountCurrency"`
	AccountID       string  `json:"AccountID"`
	Amount          float64 `json:"Amount"`
	AssetType       string  `json:"AssetType"`
	Cost            struct {
		Long struct {
			BuySell     string `json:"BuySell"`
			Currency    string `json:"Currency"`
			HoldingCost struct {
				Tax []struct {
					Pct  float64 `json:"Pct"`
					Rule struct {
						CalculationType string  `json:"CalculationType"`
						Description     string  `json:"Description"`
						Value           float64 `json:"Value"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"Tax"`
				TomNext struct {
					Pct  float64 `json:"Pct"`
					Rule struct {
						Markup float64 `json:"Markup"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"TomNext"`
			} `json:"HoldingCost"`
			TotalCost    float64 `json:"TotalCost"`
			TotalCostPct float64 `json:"TotalCostPct"`
			TradingCost  struct {
				Commissions []struct {
					Pct  float64 `json:"Pct"`
					Rule struct {
						Currency      string  `json:"Currency"`
						MinCommission float64 `json:"MinCommission"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"Commissions"`
				Spread struct {
					DisplayDecimals float64 `json:"DisplayDecimals"`
					Rule            struct {
						Pct float64 `json:"Pct"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"Spread"`
			} `json:"TradingCost"`
			TrailingCommission struct {
				Pct  float64 `json:"Pct"`
				Rule struct {
					Pct float64 `json:"Pct"`
				} `json:"Rule"`
				Value float64 `json:"Value"`
			} `json:"TrailingCommission"`
		} `json:"Long"`
		Short struct {
			BuySell     string `json:"BuySell"`
			Currency    string `json:"Currency"`
			HoldingCost struct {
				Tax []struct {
					Pct  float64 `json:"Pct"`
					Rule struct {
						CalculationType string  `json:"CalculationType"`
						Description     string  `json:"Description"`
						Value           float64 `json:"Value"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"Tax"`
				TomNext struct {
					Pct  float64 `json:"Pct"`
					Rule struct {
						Markup float64 `json:"Markup"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"TomNext"`
			} `json:"HoldingCost"`
			TotalCost    float64 `json:"TotalCost"`
			TotalCostPct float64 `json:"TotalCostPct"`
			TradingCost  struct {
				Commissions []struct {
					Pct  float64 `json:"Pct"`
					Rule struct {
						Currency      string  `json:"Currency"`
						MinCommission float64 `json:"MinCommission"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"Commissions"`
				Spread struct {
					DisplayDecimals float64 `json:"DisplayDecimals"`
					Rule            struct {
						Pct float64 `json:"Pct"`
					} `json:"Rule"`
					Value float64 `json:"Value"`
				} `json:"Spread"`
			} `json:"TradingCost"`
			TrailingCommission struct {
				Pct  float64 `json:"Pct"`
				Rule struct {
					Pct float64 `json:"Pct"`
				} `json:"Rule"`
				Value float64 `json:"Value"`
			} `json:"TrailingCommission"`
		} `json:"Short"`
	} `json:"Cost"`
	CostCalculationAssumptions []string `json:"CostCalculationAssumptions"`
	DisplayAndFormat           struct {
		BarrierFormat string  `json:"BarrierFormat"`
		Currency      string  `json:"Currency"`
		Decimals      float64 `json:"Decimals"`
		Description   string  `json:"Description"`
		Format        string  `json:"Format"`
		OrderDecimals float64 `json:"OrderDecimals"`
		StrikeFormat  string  `json:"StrikeFormat"`
		Symbol        string  `json:"Symbol"`
	} `json:"DisplayAndFormat"`
	HoldingPeriodInDays float64 `json:"HoldingPeriodInDays"`
	Instrument          string  `json:"Instrument"`
	Price               float64 `json:"Price"`
	Uic                 float64 `json:"Uic"`
}
