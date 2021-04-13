package root

type UserInfoResponse struct {
	AccessRights struct {
		CanManageCashTransfers bool `json:"CanManageCashTransfers"`
		CanTakePriceSession    bool `json:"CanTakePriceSession"`
		CanTakeTradeSession    bool `json:"CanTakeTradeSession"`
		CanTrade               bool `json:"CanTrade"`
		CanViewAnyClient       bool `json:"CanViewAnyClient"`
	} `json:"AccessRights"`
	ClientID   float64  `json:"ClientId"`
	EmployeeID string   `json:"EmployeeId"`
	Roles      []string `json:"Roles"`
	UserID     float64  `json:"UserId"`
}
