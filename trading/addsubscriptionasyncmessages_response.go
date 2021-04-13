package trading

type AddSubscriptionAsyncMessagesResponse struct {
	ContextID         string  `json:"ContextId"`
	Format            string  `json:"Format"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Snapshot          struct {
		Data []struct {
			DateTime      string `json:"DateTime"`
			DisplayType   string `json:"DisplayType"`
			IsDiscardable bool   `json:"IsDiscardable"`
			MessageBody   string `json:"MessageBody"`
			MessageHeader string `json:"MessageHeader"`
			MessageID     string `json:"MessageId"`
			MessageType   string `json:"MessageType"`
		} `json:"Data"`
	} `json:"Snapshot"`
	State string `json:"State"`
	Tag   string `json:"Tag"`
}
