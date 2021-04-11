package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/currencies/getcurrencies/1f05ea6284a3c690ec54a26051e3d055
func GetCurrencies() (*GetCurrenciesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/currencies")
	if err != nil {
		return nil, err
	}
	var respJson *GetCurrenciesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
