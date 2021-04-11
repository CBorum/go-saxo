package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/clients/getclient/1499e70934cb99a0c9e70d53f9ad8f7d
func GetClientMe() (*GetClientMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/clients/me")
	if err != nil {
		return nil, err
	}
	var respJson *GetClientMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/clients/getclient/b5e69bf214172803e7f5cd362f19f8f1
func GetClient(clientkey string) (*GetClientResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/clients/{ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClientResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/clients/updateclientsettings/be26c70709fb985d342d31aa4f5335f7
func UpdateClientSettings(params UpdateClientSettingsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/clients/me")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateClientSettingsParams struct {
	AccountValueProtectionLimit float64
	ForceOpenDefaultValue       bool
	NewPositionNettingMode      string
	NewPositionNettingProfile   string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/clients/getclients/2236c8b221db10b086f5b0d73c62c8d4
func GetClients(params GetClientsParams) (*GetClientsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/clients/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}&OwnerKey={OwnerKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClientsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClientsParams struct {
	inlinecount string
	skip        int64
	top         int64
	OwnerKey    string // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/clients/updateclientsettingsforpartner/a5b5642f959f80fa7a123aedfde53036
func UpdateClientSettingsForPartner(params UpdateClientSettingsForPartnerParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/clients/?ClientKey={ClientKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateClientSettingsForPartnerParams struct {
	AccountValueProtectionLimit float64
	ClientKey                   string // required
	ForceOpenDefaultValue       bool
	NewPositionNettingMode      string
	NewPositionNettingProfile   string
}
