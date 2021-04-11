package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/standarddates/getforwardtenordates/f3b91ed5e08e1d663c7a33bf045df567
func GetForwardTenorDates(uic string, params GetForwardTenorDatesParams) (*GetForwardTenorDatesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/standarddates/forwardtenor/{Uic}/?AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetForwardTenorDatesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetForwardTenorDatesParams struct {
	AccountKey string
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/standarddates/getfxoptionexpirydates/be7b3c6935a397693069a0529b2c7490
func GetFxOptionExpiryDates(uic string) (*GetFxOptionExpiryDatesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/standarddates/fxoptionexpiry/{Uic}")
	if err != nil {
		return nil, err
	}
	var respJson *GetFxOptionExpiryDatesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
