package autotrading

type GetInvestmentsummariesAsyncResponse struct {
	Data []struct {
		CategoryName string `json:"CategoryName"`
		Currency     string `json:"Currency"`
		Investments  []struct {
			AccountID               string  `json:"AccountId"`
			ClientID                float64 `json:"ClientId"`
			InitialFunding          float64 `json:"InitialFunding"`
			InvestmentStartDateTime string  `json:"InvestmentStartDateTime"`
			IsActive                bool    `json:"IsActive"`
			PendingFunding          float64 `json:"PendingFunding"`
			StateName               string  `json:"StateName"`
			SuitabilityLevel        string  `json:"SuitabilityLevel"`
			TimeWhenStateModified   string  `json:"TimeWhenStateModified"`
		} `json:"Investments"`
		MinimumFundingAmount float64 `json:"MinimumFundingAmount"`
		MostRecentStart      struct {
			AccountID             string  `json:"AccountId"`
			ClientID              float64 `json:"ClientId"`
			TimeWhenStateModified string  `json:"TimeWhenStateModified"`
		} `json:"MostRecentStart"`
		NumberOfFollowers  float64 `json:"NumberOfFollowers"`
		Products           string  `json:"Products"`
		ReturnPercentage   float64 `json:"ReturnPercentage"`
		StateOfLeader      string  `json:"StateOfLeader"`
		StrategyName       string  `json:"StrategyName"`
		SuitabilityLevel   string  `json:"SuitabilityLevel"`
		TradeLeaderID      string  `json:"TradeLeaderId"`
		TradeLeaderOptions string  `json:"TradeLeaderOptions"`
	} `json:"Data"`
}
