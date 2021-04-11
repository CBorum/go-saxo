package trading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v1/messages/getmessagesasync/c4548e4add65671bb9612cc152f8298c
func GetMessagesAsync() (*GetMessagesAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/trade/v1/messages")
	if err != nil {
		return nil, err
	}
	var respJson *GetMessagesAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/messages/markmessageasseen/4e74666015b9d3478088a57d85dcbc60
func MarkMessageAsSeen(messageid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/trade/v1/messages/seen/{MessageId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/messages/markmessagesasseen/8c647d4c36a23c0f9150522f57dd45f4
func MarkMessagesAsSeen(params MarkMessagesAsSeenParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/trade/v1/messages/seen/?MessageIds={MessageIds}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type MarkMessagesAsSeenParams struct {
	MessageIds string // required
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/messages/addsubscriptionasync/22ce70f29a57da59087701912c8077d1
func AddSubscriptionAsyncMessages(params AddSubscriptionAsyncMessagesParams) (*AddSubscriptionAsyncMessagesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/messages/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionAsyncMessagesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionAsyncMessagesParams struct {
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/messages/deletesubscription/b132a38a8d61746d1d4e347813d9ea79
func DeleteSubscriptionMessages(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/messages/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/messages/deletesubscriptions/d7a71b0513d2d1d99f4906e1b133387b
func DeleteSubscriptionsMessages(contextid string, params DeleteSubscriptionsMessagesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/messages/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsMessagesParams struct {
	Tag string
}
