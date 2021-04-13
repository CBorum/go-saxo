// This file is autogenerated by cmd/generate/main.go
package referencedata

import (
    "fmt"

	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/currencies/getcurrencies/1f05ea6284a3c690ec54a26051e3d055
func GetCurrencies() (*GetCurrenciesResponse, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/ref/v1/currencies", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &GetCurrenciesResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}
