package root

type AddSubscriptionResponse struct {
	ContextID         string  `json:"ContextId"`
	Format            string  `json:"Format"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Snapshot          struct {
		DataLevel  string `json:"DataLevel"`
		TradeLevel string `json:"TradeLevel"`
	} `json:"Snapshot"`
	State string `json:"State"`
}
