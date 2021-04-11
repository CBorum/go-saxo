package accounthistory

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/hist/v3/accountvalues/getstandardperiodaccountvalues/c7fad35632375ae303328622f2cf2034
func GetStandardPeriodAccountValues(clientkey string, params GetStandardPeriodAccountValuesParams) (*GetStandardPeriodAccountValuesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/hist/v3/accountvalues/{ClientKey}/?MockDataId={MockDataId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetStandardPeriodAccountValuesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetStandardPeriodAccountValuesParams struct {
	MockDataId string
}
