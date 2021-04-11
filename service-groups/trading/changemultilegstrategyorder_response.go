package trading

type ChangeMultiLegStrategyOrderResponse struct {
	MultiLegAmount  float64 `json:"MultiLegAmount"`
	MultiLegOrderID string  `json:"MultiLegOrderId"`
}
