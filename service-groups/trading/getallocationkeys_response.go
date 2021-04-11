package trading

type GetAllocationKeysResponse struct {
	Data []struct {
		AllocationKeyID   string  `json:"AllocationKeyId"`
		AllocationKeyName string  `json:"AllocationKeyName"`
		AssetType         string  `json:"AssetType"`
		BuySell           string  `json:"BuySell"`
		ExternalReference string  `json:"ExternalReference"`
		OwnerAccountKey   string  `json:"OwnerAccountKey"`
		Participants      float64 `json:"Participants"`
		Status            string  `json:"Status"`
		Uic               float64 `json:"Uic"`
	} `json:"Data"`
}
