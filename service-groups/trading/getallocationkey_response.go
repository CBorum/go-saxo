package trading

type GetAllocationKeyResponse struct {
	AllocationKeyID           string  `json:"AllocationKeyId"`
	AllocationKeyName         string  `json:"AllocationKeyName"`
	AllocationUnitType        string  `json:"AllocationUnitType"`
	AssetType                 string  `json:"AssetType"`
	BuySell                   string  `json:"BuySell"`
	ExternalReference         string  `json:"ExternalReference"`
	MarginHandling            string  `json:"MarginHandling"`
	OwnerAccountKey           string  `json:"OwnerAccountKey"`
	Participants              float64 `json:"Participants"`
	ParticipatingAccountsInfo []struct {
		AcceptRemainderAmount bool    `json:"AcceptRemainderAmount"`
		AccountKey            string  `json:"AccountKey"`
		Priority              float64 `json:"Priority"`
		UnitValue             float64 `json:"UnitValue"`
	} `json:"ParticipatingAccountsInfo"`
	Status string  `json:"Status"`
	Uic    float64 `json:"Uic"`
}
