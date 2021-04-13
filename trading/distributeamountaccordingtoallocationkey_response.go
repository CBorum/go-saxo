package trading

type DistributeAmountAccordingToAllocationKeyResponse struct {
	Remainder      float64 `json:"Remainder"`
	SubClientsInfo []struct {
		AccountID      string  `json:"AccountId"`
		Amount         float64 `json:"Amount"`
		MultiLegAmount float64 `json:"MultiLegAmount"`
	} `json:"SubClientsInfo"`
}
