package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/historicalreportdata-bookings/get/29d0de16ec94dac98d5253d314cd14bc
func GetBookings(clientkey string, params GetBookingsParams) (*GetBookingsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/reports/bookings/{ClientKey}/?$top={$top}&$skip={$skip}&$skiptoken={$skiptoken}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&FilterType={FilterType}&FilterValue={FilterValue}&FromDate={FromDate}&ToDate={ToDate}&MockDataId={MockDataId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetBookingsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetBookingsParams struct {
	skip            int64
	skiptoken       string
	top             int64
	AccountGroupKey string
	AccountKey      string

	FilterType  string
	FilterValue string
	FromDate    string
	MockDataId  string
	ToDate      string
}
