package assettransfers

type GetBeneficiaryInstructionsAsyncResponse struct {
	Data []struct {
		BeneficiaryDetails struct {
			AccountNumber string `json:"AccountNumber"`
			Bic           string `json:"Bic"`
			ClearingCode  string `json:"ClearingCode"`
			CountryCode   string `json:"CountryCode"`
		} `json:"BeneficiaryDetails"`
		BeneficiaryInstructionID string `json:"BeneficiaryInstructionId"`
		ClientID                 string `json:"ClientId"`
		Currency                 string `json:"Currency"`
		IntermediaryBank         struct {
			AccountNumber string `json:"AccountNumber"`
			Bic           string `json:"Bic"`
		} `json:"IntermediaryBank"`
		Name            string `json:"Name"`
		RegulatedBroker struct {
			AccountNumber string `json:"AccountNumber"`
			Name          string `json:"Name"`
		} `json:"RegulatedBroker"`
	} `json:"Data"`
	Count float64 `json:"__count"`
}
