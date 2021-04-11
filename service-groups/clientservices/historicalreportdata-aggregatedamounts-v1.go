package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/historicalreportdata-aggregatedamounts/get/3962da8f9c349b9c3646548a669c169e
func GetAggregatedAmounts(clientkey string, fromdate string, todate string, params GetAggregatedAmountsParams) (*GetAggregatedAmountsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/reports/aggregatedAmounts/{ClientKey}/{FromDate}/{ToDate}/?$top={$top}&$skip={$skip}&$skiptoken={$skiptoken}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&MockDataId={MockDataId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAggregatedAmountsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAggregatedAmountsParams struct {
	skip            int64
	skiptoken       string
	top             int64
	AccountGroupKey string
	AccountKey      string

	MockDataId string
}
