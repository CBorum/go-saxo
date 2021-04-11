package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/exchanges/getallexchanges/ad46a2d6611b9194c5fa84d21287dc7f
func GetAllExchanges(params GetAllExchangesParams) (*GetAllExchangesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/exchanges/?$top={$top}&$skip={$skip}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAllExchangesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAllExchangesParams struct {
	skip int64
	top  int64
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/exchanges/getexchange/3a329c212a4fbd2320f582aa50d1b772
func GetExchange(exchangeid string) (*GetExchangeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/exchanges/{ExchangeId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetExchangeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
