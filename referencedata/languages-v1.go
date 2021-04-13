// This file is autogenerated by cmd/generate/main.go
package referencedata

import (
    "fmt"

	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/languages/get/f91a16539dd69368540c9216cb519d27
func GetLanguages() (*GetLanguagesResponse, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/ref/v1/languages", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &GetLanguagesResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

