package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/countries/get/42d9cba308842cabfb9deefa2a5b3d66
func GetCountries() (*GetCountriesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/countries")
	if err != nil {
		return nil, err
	}
	var respJson *GetCountriesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
