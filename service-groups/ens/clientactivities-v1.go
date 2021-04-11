package ens

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ens/v1/clientactivities/addsubscriptionasync/0d98475e6ff38b2de76b66426008fd4a
func AddSubscriptionAsync(params AddSubscriptionAsyncParams) (*AddSubscriptionAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/ens/v1/activities/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionAsyncParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/ens/v1/clientactivities/deletesubscription/650d599c8e65c2fee9c3884ff66af1ba
func DeleteSubscription(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/ens/v1/activities/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/ens/v1/clientactivities/deletesubscriptions/b413ff7923a8fef648420ac76f1b1a86
func DeleteSubscriptions(contextid string, params DeleteSubscriptionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/ens/v1/activities/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/ens/v1/clientactivities/getactivitiesasync/873a0cb82bf75890b4929f38e2de2b6b
func GetActivitiesAsync(params GetActivitiesAsyncParams) (*GetActivitiesAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ens/v1/activities/?$top={$top}&$skiptoken={$skiptoken}&OrderStatuses={OrderStatuses}&OrderTypes={OrderTypes}&Duration={Duration}&ExpirationDateTime={ExpirationDateTime}&ClientKey={ClientKey}&AccountKey={AccountKey}&AccountGroupKey={AccountGroupKey}&IncludeSubAccounts={IncludeSubAccounts}&FromDateTime={FromDateTime}&ToDateTime={ToDateTime}&Activities={Activities}&SequenceId={SequenceId}&FieldGroups={FieldGroups}&PositionEventFilter={PositionEventFilter}")
	if err != nil {
		return nil, err
	}
	var respJson *GetActivitiesAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetActivitiesAsyncParams struct {
	skiptoken           string
	top                 int64
	AccountGroupKey     string
	AccountKey          string
	Activities          string // required
	ClientKey           string
	Duration            string
	ExpirationDateTime  string
	FieldGroups         string
	FromDateTime        string
	IncludeSubAccounts  bool
	OrderStatuses       string
	OrderTypes          string
	PositionEventFilter string
	SequenceId          string
	ToDateTime          string
}
