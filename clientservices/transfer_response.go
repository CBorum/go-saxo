package clientservices

type TransferResponse struct {
	Commission          float64 `json:"Commission"`
	FromAccountAmount   float64 `json:"FromAccountAmount"`
	FromAccountCurrency string  `json:"FromAccountCurrency"`
	ToAccountAmount     float64 `json:"ToAccountAmount"`
	ToAccountCurrency   string  `json:"ToAccountCurrency"`
}
