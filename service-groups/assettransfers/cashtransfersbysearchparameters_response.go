package assettransfers

type CashTransfersBySearchParametersResponse struct {
	Data []struct {
		Amount                float64 `json:"Amount"`
		Comment               string  `json:"Comment"`
		Currency              string  `json:"Currency"`
		ExternalReference     string  `json:"ExternalReference"`
		FromAccount           string  `json:"FromAccount"`
		FromAccountPositionID string  `json:"FromAccountPositionId"`
		FundingCheck          string  `json:"FundingCheck"`
		ToAccount             string  `json:"ToAccount"`
		ToAccountPositionID   string  `json:"ToAccountPositionId"`
		TransactionID         string  `json:"TransactionId"`
		TransferRequestedOn   string  `json:"TransferRequestedOn"`
		TransferStatus        string  `json:"TransferStatus"`
		ValueDate             string  `json:"ValueDate"`
	} `json:"Data"`
}
