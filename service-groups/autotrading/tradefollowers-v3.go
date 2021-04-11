package autotrading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradefollowers/gettradefollowers/2bdf75595f80d29308a1a78817af9555
func GetTradeFollowers() (*GetTradeFollowersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeFollowers/me")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradeFollowersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradefollowers/getsuitabilitystatus/f29fd826000313674d43a8a8f957e646
func GetSuitabilityStatus(clientkey string) (*GetSuitabilityStatusResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeFollowers/{ClientKey}/suitabilitystatus")
	if err != nil {
		return nil, err
	}
	var respJson *GetSuitabilityStatusResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradefollowers/putsuitabilitystatus/1cca3a2182825335b94faaa1b204d2eb
func PutSuitabilityStatus(clientkey string, params PutSuitabilityStatusParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/at/v3/tradeFollowers/{ClientKey}/suitabilitystatus")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type PutSuitabilityStatusParams struct {
	OverallSuitability     string // required
	ProductAreaSuitability string // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradefollowers/gettermsandconditions/1619ace40f6c5400134a4f37c50021a1
func GetTermsAndConditions(clientkey string) (*GetTermsAndConditionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeFollowers/{ClientKey}/termsandconditions")
	if err != nil {
		return nil, err
	}
	var respJson *GetTermsAndConditionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradefollowers/putclienttermsandconditions/27384a61ee7bb1390a138982f909ab54
func PutClientTermsAndConditions(clientkey string, params PutClientTermsAndConditionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/at/v3/tradeFollowers/{ClientKey}/termsandconditions")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type PutClientTermsAndConditionsParams struct {
	Accepted bool // required

}
