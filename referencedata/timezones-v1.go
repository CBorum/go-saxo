// This file is autogenerated by cmd/generate/main.go
package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/timezones/gettimezones/d1934556c178b4f09404231d2dce8233
func GetTimeZones() (*GetTimeZonesResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/ref/v1/timezones"
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetTimeZonesResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
