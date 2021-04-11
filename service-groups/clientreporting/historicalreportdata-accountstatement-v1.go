package clientreporting

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cr/v1/historicalreportdata-accountstatement/getasync/60679c07d68bb3dc5767352a81acdb16
func GetAsyncAccountStatement(clientkey string, params GetAsyncAccountStatementParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cr/v1/reports/AccountStatement/{ClientKey}/?FromDate={FromDate}&ToDate={ToDate}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GetAsyncAccountStatementParams struct {
	AccountGroupKey string
	AccountKey      string

	FromDate string
	ToDate   string
}
