package trading

type GetMessagesAsyncResponse struct {
	Data []struct {
		DateTime      string `json:"DateTime"`
		DisplayType   string `json:"DisplayType"`
		IsDiscardable bool   `json:"IsDiscardable"`
		MessageBody   string `json:"MessageBody"`
		MessageHeader string `json:"MessageHeader"`
		MessageID     string `json:"MessageId"`
		MessageType   string `json:"MessageType"`
	} `json:"Data"`
}
