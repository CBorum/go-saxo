package autotrading

type GetInvestmentResponse struct {
	AccountID              string  `json:"AccountId"`
	AccountKey             string  `json:"AccountKey"`
	ClientID               float64 `json:"ClientId"`
	Currency               string  `json:"Currency"`
	DisplayName            string  `json:"DisplayName"`
	DisplayState           float64 `json:"DisplayState"`
	EntryGateResult        string  `json:"EntryGateResult"`
	ErrorNumber            float64 `json:"ErrorNumber"`
	InitialFunding         float64 `json:"InitialFunding"`
	InvestmentID           string  `json:"InvestmentId"`
	InvestmentProcessState string  `json:"InvestmentProcessState"`
	InvestmentShieldAmount float64 `json:"InvestmentShieldAmount"`
	InvestmentStateID      float64 `json:"InvestmentStateId"`
	IsAuthorizedToFollow   bool    `json:"IsAuthorizedToFollow"`
	IsFollowAllowed        bool    `json:"IsFollowAllowed"`
	IsFollowing            bool    `json:"IsFollowing"`
	IsOpenForFollowers     bool    `json:"IsOpenForFollowers"`
	IsReadyForTrading      bool    `json:"IsReadyForTrading"`
	IsTradeFollowerReady   bool    `json:"IsTradeFollowerReady"`
	IsWithdrawalInProgress bool    `json:"IsWithdrawalInProgress"`
	MinimumFunding         float64 `json:"MinimumFunding"`
	PendingFunding         float64 `json:"PendingFunding"`
	ReservedAmount         float64 `json:"ReservedAmount"`
	StrategyName           string  `json:"StrategyName"`
	TradeLeaderID          string  `json:"TradeLeaderId"`
}
