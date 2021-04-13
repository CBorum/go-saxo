package trading

type CancelAllOrderResponse struct {
	ErrorInfo struct {
		ErrorCode string `json:"ErrorCode"`
		Message   string `json:"Message"`
	} `json:"ErrorInfo"`
}
