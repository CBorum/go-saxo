package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/getopenorder/3024791d13ad168eeb5729c6f65659a3
func GetOpenOrder(clientkey string, orderid string, params GetOpenOrderParams) (*GetOpenOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/orders/{ClientKey}/{OrderId}/?FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetOpenOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetOpenOrderParams struct {
	FieldGroups string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/getopenorders/cae24349737ea6fef4f0e6d9c477794e
func GetOpenOrdersMe(params GetOpenOrdersMeParams) (*GetOpenOrdersMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/orders/me/?$top={$top}&$skip={$skip}&FieldGroups={FieldGroups}&Status={Status}&MultiLegOrderId={MultiLegOrderId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetOpenOrdersMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetOpenOrdersMeParams struct {
	skip            int64
	top             int64
	FieldGroups     string
	MultiLegOrderId string
	Status          string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/getopenorderdetails/4344116e34c37b70b5f7b8394f406934
func GetOpenOrderDetails(orderid string, params GetOpenOrderDetailsParams) (*GetOpenOrderDetailsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/orders/{OrderId}/details/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetOpenOrderDetailsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetOpenOrderDetailsParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required

}

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/getopenorders/aef7f2c761e8f985b3cf02372d3a21be
func GetOpenOrders(params GetOpenOrdersParams) (*GetOpenOrdersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/orders/?$top={$top}&$skip={$skip}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&OrderId={OrderId}&Status={Status}&FieldGroups={FieldGroups}&WatchlistId={WatchlistId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetOpenOrdersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetOpenOrdersParams struct {
	skip            int64
	top             int64
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
	FieldGroups     string
	OrderId         string
	Status          string
	WatchlistId     string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/addactivesubscription/0de7616778e992af51f475849270bc2e
func AddActiveSubscriptionOrders(params AddActiveSubscriptionOrdersParams) (*AddActiveSubscriptionOrdersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/orders/subscriptions/?$top={$top}&$skip={$skip}&$skiptoken={$skiptoken}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *AddActiveSubscriptionOrdersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddActiveSubscriptionOrdersParams struct {
	inlinecount string
	skip        int64
	skiptoken   string
	top         int64
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/deletesubscriptions/1ed8d960884aac2fb4b89b8d207d33aa
func DeleteSubscriptionsOrders(contextid string, params DeleteSubscriptionsOrdersParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/orders/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsOrdersParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/orders/deletesubscription/27113d973d511604762399997db5b9b7
func DeleteSubscriptionOrders(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/orders/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
