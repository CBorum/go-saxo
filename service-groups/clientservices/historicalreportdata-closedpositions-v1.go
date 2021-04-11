package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/historicalreportdata-closedpositions/get/9fe036fb16df327459bc4cbfb7359cb0
func GetClosedPositions(clientkey string, fromdate string, todate string, params GetClosedPositionsParams) (*GetClosedPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/reports/closedPositions/{ClientKey}/{FromDate}/{ToDate}/?$top={$top}&$skip={$skip}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClosedPositionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClosedPositionsParams struct {
	skip            int64
	top             int64
	AccountGroupKey string
	AccountKey      string
}
