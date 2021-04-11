package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/audit-orderactivities/getorderstatesasync/88396c9accc21e373925b5cd2ce134dd
func GetOrderStatesAsync(params GetOrderStatesAsyncParams) (*GetOrderStatesAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/audit/orderactivities/?$top={$top}&$skiptoken={$skiptoken}&FieldGroups={FieldGroups}&Status={Status}&ClientKey={ClientKey}&AccountKey={AccountKey}&OrderId={OrderId}&EntryType={EntryType}&FromDateTime={FromDateTime}&ToDateTime={ToDateTime}&CorrelationKey={CorrelationKey}&IncludeSubAccounts={IncludeSubAccounts}")
	if err != nil {
		return nil, err
	}
	var respJson *GetOrderStatesAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetOrderStatesAsyncParams struct {
	skiptoken          string
	top                int64
	AccountKey         string
	ClientKey          string
	CorrelationKey     string
	EntryType          string
	FieldGroups        string
	FromDateTime       string
	IncludeSubAccounts bool
	OrderId            string
	Status             string
	ToDateTime         string
}
