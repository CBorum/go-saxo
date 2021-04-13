package assettransfers

type GetTransferDetailsResponse struct {
	Data []struct {
		AccountID      string `json:"AccountId"`
		AccountManager string `json:"AccountManager"`
		AssetType      string `json:"AssetType"`
		Broker         struct {
			Contact     string `json:"Contact"`
			CountryCode string `json:"CountryCode"`
			Email       string `json:"Email"`
			Name        string `json:"Name"`
			Phone       string `json:"Phone"`
		} `json:"Broker"`
		ClientAccountAtBroker string `json:"ClientAccountAtBroker"`
		Securities            []struct {
			Currency             string  `json:"Currency"`
			Description          string  `json:"Description"`
			Exchange             string  `json:"Exchange"`
			InstrumentSymbolCode string  `json:"InstrumentSymbolCode"`
			IsinCode             string  `json:"IsinCode"`
			Price                float64 `json:"Price"`
			PurchaseDate         string  `json:"PurchaseDate"`
			Quantity             float64 `json:"Quantity"`
			Uic                  float64 `json:"Uic"`
		} `json:"Securities"`
		TradeDate         string `json:"TradeDate"`
		TransferDirection string `json:"TransferDirection"`
		TransferID        string `json:"TransferId"`
	} `json:"Data"`
	Count float64 `json:"__count"`
}
