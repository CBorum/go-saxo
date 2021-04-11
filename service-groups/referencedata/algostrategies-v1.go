package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/algostrategies/getallstrategies/fce5d6b477a3e4d2a1d83e609f02d24f
func GetAllStrategies(params GetAllStrategiesParams) (*GetAllStrategiesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/algostrategies/?$top={$top}&$skip={$skip}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAllStrategiesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAllStrategiesParams struct {
	skip int64
	top  int64
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/algostrategies/getstrategybyname/f3195d507d54dd21c5504df61bf17863
func GetStrategyByName(name string) (*GetStrategyByNameResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/algostrategies/{Name}")
	if err != nil {
		return nil, err
	}
	var respJson *GetStrategyByNameResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
