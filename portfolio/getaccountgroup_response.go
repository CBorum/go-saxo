package portfolio

type GetAccountGroupResponse struct {
	AccountGroupKey                     string  `json:"AccountGroupKey"`
	AccountGroupName                    string  `json:"AccountGroupName"`
	AccountValueProtectionLimit         float64 `json:"AccountValueProtectionLimit"`
	CollateralMonitoringMode            string  `json:"CollateralMonitoringMode"`
	MarginCalculationMethod             string  `json:"MarginCalculationMethod"`
	MarginMonitoringMode                string  `json:"MarginMonitoringMode"`
	SupportsAccountValueProtectionLimit bool    `json:"SupportsAccountValueProtectionLimit"`
}
