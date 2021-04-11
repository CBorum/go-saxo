package portfolio

type GetUsersResponse struct {
	Data []struct {
		ClientKey                         string   `json:"ClientKey"`
		Culture                           string   `json:"Culture"`
		Language                          string   `json:"Language"`
		LastLoginStatus                   string   `json:"LastLoginStatus"`
		LastLoginTime                     string   `json:"LastLoginTime"`
		LegalAssetTypes                   []string `json:"LegalAssetTypes"`
		MarketDataViaOpenAPITermsAccepted bool     `json:"MarketDataViaOpenApiTermsAccepted"`
		Name                              string   `json:"Name"`
		TimeZoneID                        float64  `json:"TimeZoneId"`
		UserID                            string   `json:"UserId"`
		UserKey                           string   `json:"UserKey"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
