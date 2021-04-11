package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/timezones/gettimezones/d1934556c178b4f09404231d2dce8233
func GetTimeZones() (*GetTimeZonesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/timezones")
	if err != nil {
		return nil, err
	}
	var respJson *GetTimeZonesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
