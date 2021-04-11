package trading

type CancelOrderResponse struct {
	Orders []struct {
		OrderID string `json:"OrderId"`
	} `json:"Orders"`
}
