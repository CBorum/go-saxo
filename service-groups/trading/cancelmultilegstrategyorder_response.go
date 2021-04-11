package trading

type CancelMultiLegStrategyOrderResponse struct {
	Orders []struct {
		OrderID string `json:"OrderId"`
	} `json:"Orders"`
}
