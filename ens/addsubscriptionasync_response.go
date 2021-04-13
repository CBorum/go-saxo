package ens

type AddSubscriptionAsyncResponse struct {
	ContextID         string  `json:"ContextId"`
	InactivityTimeout float64 `json:"InactivityTimeout"`
	ReferenceID       string  `json:"ReferenceId"`
	RefreshRate       float64 `json:"RefreshRate"`
	Snapshot          struct {
		Data []interface{} `json:"Data"`
	} `json:"Snapshot"`
	State string `json:"State"`
}
