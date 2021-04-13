package clientservices

type SearchResponse struct {
	Data []struct {
		Accounts []struct {
			AccountDisplayName string `json:"AccountDisplayName"`
			AccountGroupID     string `json:"AccountGroupId"`
			AccountGroupName   string `json:"AccountGroupName"`
			AccountID          string `json:"AccountId"`
			AccountKey         string `json:"AccountKey"`
		} `json:"Accounts"`
		ClientID          string   `json:"ClientId"`
		ClientKey         string   `json:"ClientKey"`
		ClientType        string   `json:"ClientType"`
		ContractType      string   `json:"ContractType"`
		DefaultAccountID  string   `json:"DefaultAccountId"`
		DefaultAccountKey string   `json:"DefaultAccountKey"`
		LegalAssetTypes   []string `json:"LegalAssetTypes"`
		Name              string   `json:"Name"`
		PartnerPlatformID string   `json:"PartnerPlatformId"`
		Users             []struct {
			UserID  string `json:"UserId"`
			UserKey string `json:"UserKey"`
		} `json:"Users"`
	} `json:"Data"`
}
