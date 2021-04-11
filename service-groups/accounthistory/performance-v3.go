package accounthistory

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/hist/v3/performance/getperformance/9dd399b9b4d92ce3876ee9e787ded861
func GetPerformance(clientkey string, params GetPerformanceParams) (*GetPerformanceResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/hist/v3/perf/{ClientKey}/?AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&FromDate={FromDate}&ToDate={ToDate}&StandardPeriod={StandardPeriod}&MockDataId={MockDataId}&FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetPerformanceResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetPerformanceParams struct {
	AccountGroupKey string
	AccountKey      string

	FieldGroups    string
	FromDate       string
	MockDataId     string
	StandardPeriod string
	ToDate         string
}
