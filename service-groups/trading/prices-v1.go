package trading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/addsubscriptionasync/966e8dd4a6c8a1d9a210200faaf0a1e3
func AddSubscriptionAsyncPrices(params AddSubscriptionAsyncPricesParams) (*AddSubscriptionAsyncPricesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionAsyncPricesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionAsyncPricesParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/requestmarginimpactonsubscription/8ba2bfe2b9cb41a0b1af66f8e26569ac
func RequestMarginImpactOnSubscription(contextid string, referenceid string) (*RequestMarginImpactOnSubscriptionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/subscriptions/{ContextId}/{ReferenceId}/MarginImpact")
	if err != nil {
		return nil, err
	}
	var respJson *RequestMarginImpactOnSubscriptionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/deletesubscriptions/7bde8ba8002b53fad611872074bc351d
func DeleteSubscriptionsPrices(contextid string, params DeleteSubscriptionsPricesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsPricesParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/deletesubscription/1d74d35f6ac026bf0e8043eabaaa5c00
func DeleteSubscriptionPrices(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/getmultilegpriceasync/58d97a64e4d0cb8b899e228b658ddc2e
func GetMultiLegPriceAsync(params GetMultiLegPriceAsyncParams) (*GetMultiLegPriceAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/multileg")
	if err != nil {
		return nil, err
	}
	var respJson *GetMultiLegPriceAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetMultiLegPriceAsyncParams struct {
	AccountKey  string
	FieldGroups string
	Legs        string // required
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/addmultilegpricessubscriptionasync/181b016c8d910b195e9ab62cb8f45f60
func AddMultiLegPricesSubscriptionAsync(params AddMultiLegPricesSubscriptionAsyncParams) (*AddMultiLegPricesSubscriptionAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/multileg/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddMultiLegPricesSubscriptionAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddMultiLegPricesSubscriptionAsyncParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/deletemultilegpricessubscriptions/376fe388c6b7db174a735f97b360b364
func DeleteMultiLegPricesSubscriptions(contextid string, params DeleteMultiLegPricesSubscriptionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/multileg/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteMultiLegPricesSubscriptionsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/prices/deletemultilegpricessubscription/ce145ca22d78a50b10b715b5139ba156
func DeleteMultiLegPricesSubscription(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/prices/multileg/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
