package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/historicalreportdata-trades/getdetailsasync/530c3d12f21817d3a77911a4644fa0f1
func GetDetailsAsync(clientkey string, params GetDetailsAsyncParams) (*GetDetailsAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/reports/trades/{ClientKey}/?$top={$top}&$skip={$skip}&$skiptoken={$skiptoken}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&FromDate={FromDate}&ToDate={ToDate}&TradeId={TradeId}&MockDataId={MockDataId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetDetailsAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetDetailsAsyncParams struct {
	skip            int64
	skiptoken       string
	top             int64
	AccountGroupKey string
	AccountKey      string

	FromDate   string
	MockDataId string
	ToDate     string
	TradeId    string
}
