package portfolio

type GetAllClientEntitlementsResponse struct {
	Data []struct {
		Entitlements []struct {
			Greeks            []string `json:"Greeks"`
			RealTimeFullBook  []string `json:"RealTimeFullBook"`
			RealTimeTopOfBook []string `json:"RealTimeTopOfBook"`
		} `json:"Entitlements"`
		ExchangeID string `json:"ExchangeId"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
