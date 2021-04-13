package clientservices

type GetWireTransfersResponse struct {
	ClientDetails struct {
		AccountID            string  `json:"AccountId"`
		AccountNumber        string  `json:"AccountNumber"`
		AttainKey            string  `json:"AttainKey"`
		Bic                  string  `json:"Bic"`
		ClientID             float64 `json:"ClientId"`
		ClientNameAndAddress string  `json:"ClientNameAndAddress"`
		Iban                 string  `json:"Iban"`
		RegistrationNumber   string  `json:"RegistrationNumber"`
	} `json:"ClientDetails"`
	CurrencyCode     string `json:"CurrencyCode"`
	IntermediaryBank struct {
		Address string `json:"Address"`
		Bic     string `json:"Bic"`
		Name    string `json:"Name"`
	} `json:"IntermediaryBank"`
	IsBrokerFunding bool `json:"IsBrokerFunding"`
	ReceivingBank   struct {
		Address string `json:"Address"`
		Bic     string `json:"Bic"`
		Name    string `json:"Name"`
	} `json:"ReceivingBank"`
	SubsidiaryID     string `json:"SubsidiaryId"`
	SubsidiaryName   string `json:"SubsidiaryName"`
	VirtualAccountID string `json:"VirtualAccountId"`
}
