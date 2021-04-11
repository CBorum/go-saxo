package portfolio

type GetAccountGroupsForLoggedInUserResponse struct {
	Data []struct {
		AccountGroupKey                     string  `json:"AccountGroupKey"`
		AccountGroupName                    string  `json:"AccountGroupName"`
		AccountValueProtectionLimit         float64 `json:"AccountValueProtectionLimit"`
		CollateralMonitoringMode            string  `json:"CollateralMonitoringMode"`
		MarginCalculationMethod             string  `json:"MarginCalculationMethod"`
		MarginMonitoringMode                string  `json:"MarginMonitoringMode"`
		SupportsAccountValueProtectionLimit bool    `json:"SupportsAccountValueProtectionLimit"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
