package trading

type PlaceMultiLegStrategyOrderResponse struct {
	MultiLegAmount  float64 `json:"MultiLegAmount"`
	MultiLegOrderID string  `json:"MultiLegOrderId"`
	Orders          []struct {
		OrderID string `json:"OrderId"`
	} `json:"Orders"`
}
