package portfolio

type GetMarginOverviewResponse struct {
	Groups []struct {
		Contributors []struct {
			AssetTypes            []string `json:"AssetTypes"`
			InstrumentDescription string   `json:"InstrumentDescription"`
			InstrumentSpecifier   string   `json:"InstrumentSpecifier"`
			Margin                float64  `json:"Margin"`
			Uic                   float64  `json:"Uic"`
		} `json:"Contributors"`
		GroupType   string  `json:"GroupType"`
		TotalMargin float64 `json:"TotalMargin"`
	} `json:"Groups"`
}
