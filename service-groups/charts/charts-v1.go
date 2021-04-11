package charts

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/chart/v1/charts/getchartdataasync/387cfc61d3292d9237095b9144ac4733
func GetChartDataAsync(params GetChartDataAsyncParams) (*GetChartDataAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/chart/v1/charts/?AssetType={AssetType}&Count={Count}&FieldGroups={FieldGroups}&Horizon={Horizon}&Mode={Mode}&Time={Time}&Uic={Uic}")
	if err != nil {
		return nil, err
	}
	var respJson *GetChartDataAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetChartDataAsyncParams struct {
	AssetType   string // required
	Count       int64
	FieldGroups string
	Horizon     int64 // required
	Mode        string
	Time        string
	Uic         int64 // required
}

// https://www.developer.saxo/openapi/referencedocs/chart/v1/charts/addsubscriptionasync/ffe4364c2a06b804b58ce27ad5c8d521
func AddSubscriptionAsync(params AddSubscriptionAsyncParams) (*AddSubscriptionAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/chart/v1/charts/subscriptions")
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

// https://www.developer.saxo/openapi/referencedocs/chart/v1/charts/deletesubscriptions/1f8284579f1a2d3b594c2e6855b14f63
func DeleteSubscriptions(contextid string, params DeleteSubscriptionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/chart/v1/charts/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/chart/v1/charts/deletesubscription/017501f47274e61e09d79c1bed9680df
func DeleteSubscription(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/chart/v1/charts/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
