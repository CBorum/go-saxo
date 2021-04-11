package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/languages/get/f91a16539dd69368540c9216cb519d27
func GetLanguages() (*GetLanguagesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/languages")
	if err != nil {
		return nil, err
	}
	var respJson *GetLanguagesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
