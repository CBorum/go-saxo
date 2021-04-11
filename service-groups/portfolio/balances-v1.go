package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/balances/getbalance/9f4b3d9066a318235b25616bb21a672e
func GetBalanceMe() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/balances/me")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/balances/getbalance/2cddf1f35bd59dd2f55252c26fdf351a
func GetBalance(params GetBalanceParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/balances/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GetBalanceParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
	FieldGroups     string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/balances/getmarginoverview/e9b1b0df85178358a36cd8be3dc97226
func GetMarginOverview(params GetMarginOverviewParams) (*GetMarginOverviewResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/balances/marginoverview/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetMarginOverviewResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetMarginOverviewParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/balances/addactivesubscription/aec00bc9b272be3c599e8347929f9a2a
func AddActiveSubscriptionBalances(params AddActiveSubscriptionBalancesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/balances/subscriptions")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type AddActiveSubscriptionBalancesParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/balances/deletesubscriptions/90c8f783152a8cfa904622a8b4db948d
func DeleteSubscriptionsBalances(contextid string, params DeleteSubscriptionsBalancesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/balances/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsBalancesParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/balances/deletesubscription/269cef4ca3fa9cd15f18101f722304e5
func DeleteSubscriptionBalances(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/balances/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
