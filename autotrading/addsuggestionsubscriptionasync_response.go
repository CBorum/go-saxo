package autotrading

type AddSuggestionSubscriptionAsyncResponse struct {
	ContextID         string  `json:"ContextId"`
	Format            string  `json:"Format"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Schema            string  `json:"Schema"`
	SchemaName        string  `json:"SchemaName"`
	Snapshot          []struct {
		AutoTradingPartnerLeaderID string  `json:"AutoTradingPartnerLeaderId"`
		ClientID                   float64 `json:"ClientId"`
		Currency                   string  `json:"Currency"`
		IsAuthorizedToFollow       bool    `json:"IsAuthorizedToFollow"`
		IsFollowAllowed            bool    `json:"IsFollowAllowed"`
		IsFollowing                bool    `json:"IsFollowing"`
		IsOpenForFollowers         bool    `json:"IsOpenForFollowers"`
		IsReadyForTrading          bool    `json:"IsReadyForTrading"`
		IsTradeFollowerReady       bool    `json:"IsTradeFollowerReady"`
		MinimumFunding             float64 `json:"MinimumFunding"`
		StrategyName               string  `json:"StrategyName"`
		TradeLeaderID              string  `json:"TradeLeaderId"`
	} `json:"Snapshot"`
	State string `json:"State"`
	Tag   string `json:"Tag"`
}
