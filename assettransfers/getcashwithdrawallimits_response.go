package assettransfers

type GetCashWithdrawalLimitsResponse struct {
	Data []struct {
		AccountKey                 string  `json:"AccountKey"`
		ClientKey                  string  `json:"ClientKey"`
		MaxAllowedWithdrawalAmount float64 `json:"MaxAllowedWithdrawalAmount"`
	} `json:"Data"`
}
