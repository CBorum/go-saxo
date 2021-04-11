package referencedata

type GetAllStrategiesResponse struct {
	Data []struct {
		Description  string  `json:"Description"`
		MinAmountUSD float64 `json:"MinAmountUSD"`
		Name         string  `json:"Name"`
		Parameters   []struct {
			DataType        float64 `json:"DataType"`
			IsEditable      bool    `json:"IsEditable"`
			IsMandatory     bool    `json:"IsMandatory"`
			UIDefaultValue  string  `json:"UiDefaultValue"`
			UIOrderingIndex float64 `json:"UiOrderingIndex"`
			UIStepSize      float64 `json:"UiStepSize"`
		} `json:"Parameters"`
		SupportedDurationTypes  []string `json:"SupportedDurationTypes"`
		SupportedOrderTypes     []string `json:"SupportedOrderTypes"`
		TradableInstrumentTypes []string `json:"TradableInstrumentTypes"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
