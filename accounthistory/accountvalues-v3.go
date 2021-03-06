// This file is autogenerated by cmd/generate/main.go
package accounthistory

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/hist/v3/accountvalues/getstandardperiodaccountvalues/c7fad35632375ae303328622f2cf2034
func GetStandardPeriodAccountValues(clientkey string, params *GetStandardPeriodAccountValuesParams) (*GetStandardPeriodAccountValuesResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/hist/v3/accountvalues/{ClientKey}/?MockDataId={MockDataId}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{ClientKey}", clientkey))
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetStandardPeriodAccountValuesResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetStandardPeriodAccountValuesParams struct {
	MockDataId string `url:",omitempty"`
}
