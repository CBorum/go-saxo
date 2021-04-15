// This file is autogenerated by cmd/generate/main.go
package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/currencypairs/getallcurrencypairs/ed29edb888a5d9cbb06a253547d31c62
func GetAllCurrencyPairs(params *GetAllCurrencyPairsParams) (*GetAllCurrencyPairsResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/ref/v1/currencypairs/?AccountKey={AccountKey}&ClientKey={ClientKey}"
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetAllCurrencyPairsResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAllCurrencyPairsParams struct {
	AccountKey string `url:",omitempty"`
	ClientKey  string `url:",omitempty"`
}
