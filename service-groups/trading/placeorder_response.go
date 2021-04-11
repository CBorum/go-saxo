package trading

type PlaceOrderResponse struct {
	Orders []struct {
		OrderID string `json:"OrderId"`
	} `json:"Orders"`
}
