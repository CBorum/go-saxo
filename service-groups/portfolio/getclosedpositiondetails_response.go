package portfolio

type GetClosedPositionDetailsResponse struct {
	ClosedPosition struct {
		AccountID                                    string  `json:"AccountId"`
		Amount                                       float64 `json:"Amount"`
		AssetType                                    string  `json:"AssetType"`
		BuyOrSell                                    string  `json:"BuyOrSell"`
		ClientID                                     string  `json:"ClientId"`
		ClosedProfitLoss                             float64 `json:"ClosedProfitLoss"`
		ClosedProfitLossInBaseCurrency               float64 `json:"ClosedProfitLossInBaseCurrency"`
		ClosingMarketValue                           float64 `json:"ClosingMarketValue"`
		ClosingMarketValueInBaseCurrency             float64 `json:"ClosingMarketValueInBaseCurrency"`
		ClosingMethod                                string  `json:"ClosingMethod"`
		ClosingPositionID                            string  `json:"ClosingPositionId"`
		ClosingPrice                                 float64 `json:"ClosingPrice"`
		ConversionRateInstrumentToBaseSettledClosing bool    `json:"ConversionRateInstrumentToBaseSettledClosing"`
		ConversionRateInstrumentToBaseSettledOpening bool    `json:"ConversionRateInstrumentToBaseSettledOpening"`
		ExecutionTimeClose                           string  `json:"ExecutionTimeClose"`
		ExecutionTimeOpen                            string  `json:"ExecutionTimeOpen"`
		OpenPrice                                    float64 `json:"OpenPrice"`
		OpeningPositionID                            string  `json:"OpeningPositionId"`
		Uic                                          float64 `json:"Uic"`
	} `json:"ClosedPosition"`
	ClosedPositionDetails struct {
		CurrencyConversionRateInstrumentToBaseClosing float64 `json:"CurrencyConversionRateInstrumentToBaseClosing"`
		CurrencyConversionRateInstrumentToBaseOpening float64 `json:"CurrencyConversionRateInstrumentToBaseOpening"`
		ValueDateClose                                string  `json:"ValueDateClose"`
		ValueDateOpen                                 string  `json:"ValueDateOpen"`
	} `json:"ClosedPositionDetails"`
	ClosedPositionUniqueID string `json:"ClosedPositionUniqueId"`
	DisplayAndFormat       struct {
		Currency    string  `json:"Currency"`
		Decimals    float64 `json:"Decimals"`
		Description string  `json:"Description"`
		Format      string  `json:"Format"`
		Symbol      string  `json:"Symbol"`
	} `json:"DisplayAndFormat"`
	Exchange struct {
		Description string `json:"Description"`
		ExchangeID  string `json:"ExchangeId"`
		IsOpen      bool   `json:"IsOpen"`
	} `json:"Exchange"`
	NetPositionID string `json:"NetPositionId"`
}
