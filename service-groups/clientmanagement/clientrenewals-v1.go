package clientmanagement

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cm/v1/clientrenewals/getclientrenewaldata/3ecfec627627fcb34854c255b76c8214
func GetClientRenewalData(params GetClientRenewalDataParams) (*GetClientRenewalDataResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cm/v1/clientrenewals/?ClientKey={ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClientRenewalDataResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClientRenewalDataParams struct {
	ClientKey string
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/clientrenewals/updateclientrenewaldata/3f4f8cd016fcbdcd1b39d14c009bec06
func UpdateClientRenewalData(renewalentityid string, params UpdateClientRenewalDataParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/cm/v1/clientrenewals/{RenewalEntityId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateClientRenewalDataParams struct {
	ClientKey   string // required
	Documents   string
	RenewalData string
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/clientrenewals/getrenewalstatuses/d5820d08a797c7eccacfdb1ee80fd442
func GetRenewalStatuses(params GetRenewalStatusesParams) (*GetRenewalStatusesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cm/v1/clientrenewals/pending/?$top={$top}&$skip={$skip}&OwnerKey={OwnerKey}&MustRenewBy={MustRenewBy}")
	if err != nil {
		return nil, err
	}
	var respJson *GetRenewalStatusesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetRenewalStatusesParams struct {
	skip        int64
	top         int64
	MustRenewBy string
	OwnerKey    string
}