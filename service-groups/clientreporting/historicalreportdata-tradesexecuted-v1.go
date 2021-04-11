package clientreporting

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cr/v1/historicalreportdata-tradesexecuted/getasync/b2a76e81cbe2daf0a95119118c2e701e
func GetAsyncTradesExecuted(clientkey string, params GetAsyncTradesExecutedParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cr/v1/reports/TradesExecuted/{ClientKey}/?FromDate={FromDate}&ToDate={ToDate}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GetAsyncTradesExecutedParams struct {
	AccountGroupKey string
	AccountKey      string

	FromDate string
	ToDate   string
}
