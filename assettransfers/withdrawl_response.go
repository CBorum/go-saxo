package assettransfers

type WithdrawlResponse struct {
	CutExchangeRate     float64 `json:"CutExchangeRate"`
	WithdrawalRequestID string  `json:"WithdrawalRequestId"`
}
