package trading

type AddSubscriptionAsyncOptionsChainResponse struct {
	ContextID         string  `json:"ContextId"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Snapshot          struct {
		AssetType string `json:"AssetType"`
		Expiries  []struct {
			DisplayDate            string  `json:"DisplayDate"`
			DisplayDaysToExpiry    float64 `json:"DisplayDaysToExpiry"`
			Expiry                 string  `json:"Expiry"`
			Index                  float64 `json:"Index"`
			LastTradeDate          string  `json:"LastTradeDate"`
			MidStrikePrice         float64 `json:"MidStrikePrice"`
			StrikeCount            float64 `json:"StrikeCount"`
			StrikeWindowStartIndex float64 `json:"StrikeWindowStartIndex"`
			Strikes                []struct {
				Call struct {
					High         float64 `json:"High"`
					LastClose    float64 `json:"LastClose"`
					LastTraded   float64 `json:"LastTraded"`
					Low          float64 `json:"Low"`
					NetChange    float64 `json:"NetChange"`
					Open         float64 `json:"Open"`
					OpenInterest float64 `json:"OpenInterest"`
					PriceTypeAsk string  `json:"PriceTypeAsk"`
					PriceTypeBid string  `json:"PriceTypeBid"`
					Uic          float64 `json:"Uic"`
					Volume       float64 `json:"Volume"`
				} `json:"Call"`
				Index float64 `json:"Index"`
				Put   struct {
					High         float64 `json:"High"`
					LastClose    float64 `json:"LastClose"`
					LastTraded   float64 `json:"LastTraded"`
					Low          float64 `json:"Low"`
					NetChange    float64 `json:"NetChange"`
					Open         float64 `json:"Open"`
					OpenInterest float64 `json:"OpenInterest"`
					PriceTypeAsk string  `json:"PriceTypeAsk"`
					PriceTypeBid string  `json:"PriceTypeBid"`
					Uic          float64 `json:"Uic"`
					Volume       float64 `json:"Volume"`
				} `json:"Put"`
				Strike float64 `json:"Strike"`
			} `json:"Strikes"`
			UnderlyingUic float64 `json:"UnderlyingUic"`
		} `json:"Expiries"`
		ExpiryCount float64 `json:"ExpiryCount"`
		LastUpdated string  `json:"LastUpdated"`
	} `json:"Snapshot"`
	State string `json:"State"`
}
