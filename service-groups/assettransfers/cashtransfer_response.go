package assettransfers

type CashTransferResponse struct {
	FromAccountPositionID string `json:"FromAccountPositionId"`
	ToAccountPositionID   string `json:"ToAccountPositionId"`
	TransactionID         string `json:"TransactionId"`
}
