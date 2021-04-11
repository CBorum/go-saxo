package clientreporting

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cr/v1/historicalreportdata-tradedetails/getasync/982c285ce68c79f21b74914bc6d5f6b7
func GetAsyncTradeDetails(clientkey string, params GetAsyncTradeDetailsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cr/v1/reports/TradeDetails/{ClientKey}/?AccountKey={AccountKey}&TradeId={TradeId}&FilterType={FilterType}&FilterValue={FilterValue}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GetAsyncTradeDetailsParams struct {
	AccountKey string // required

	FilterType  string // required
	FilterValue string // required
	TradeId     string // required
}
