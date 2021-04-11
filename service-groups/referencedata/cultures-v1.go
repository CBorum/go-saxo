package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/cultures/getcultures/eb75bb169352bb712a2d8c1bde115ecc
func GetCultures() (*GetCulturesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/cultures")
	if err != nil {
		return nil, err
	}
	var respJson *GetCulturesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
