package developer

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getsecrets/11e55e5f3a4e8d882e41b0120f734e3f
func GetSecrets(appkey string, params GetSecretsParams) (*GetSecretsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/secrets/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *GetSecretsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetSecretsParams struct {
	inlinecount string
	skip        int64
	top         int64
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getsecretbyid/bb1b1562bd4f9f4cfa3fca6ff73e500f
func GetSecretById(appkey string, secretid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/secrets/{SecretId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/createsecret/25e379e2acd69b5fd80781134e352228
func CreateSecret(appkey string, params CreateSecretParams) (*CreateSecretResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/secrets")
	if err != nil {
		return nil, err
	}
	var respJson *CreateSecretResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CreateSecretParams struct {
	ValidFrom  string // required
	ValidUntil string // required
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/updatesecret/7ce39fa5ac55fdcc6d26ce824b961438
func UpdateSecret(appkey string, secretid string, params UpdateSecretParams) (*UpdateSecretResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/secrets/{SecretId}")
	if err != nil {
		return nil, err
	}
	var respJson *UpdateSecretResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type UpdateSecretParams struct {
	Regenerate bool

	ValidFrom  string
	ValidUntil string
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/deletesecret/b7fee7f8a88e8c779022c5bfa662cafc
func DeleteSecret(appkey string, secretid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/secrets/{SecretId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getredirecturis/36e137507b1716182cefcc588ac8ea7b
func GetRedirectUris(appkey string, params GetRedirectUrisParams) (*GetRedirectUrisResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/redirecturis/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *GetRedirectUrisResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetRedirectUrisParams struct {
	inlinecount string
	skip        int64
	top         int64
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getredirecturifromid/79451fe3588639e19aee39d09f726aa6
func GetRedirectUriFromId(appkey string, redirecturiid string, params GetRedirectUriFromIdParams) (*GetRedirectUriFromIdResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/redirecturis/{RedirectUriId}/?$top={$top}&$skip={$skip}&$skiptoken={$skiptoken}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *GetRedirectUriFromIdResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetRedirectUriFromIdParams struct {
	inlinecount string
	skip        int64
	skiptoken   string
	top         int64
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/updateredirecturi/ed465762bf97624e41b445d3f3029ad9
func UpdateRedirectUri(appkey string, redirecturiid string, params UpdateRedirectUriParams) (*UpdateRedirectUriResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/redirecturis/{RedirectUriId}")
	if err != nil {
		return nil, err
	}
	var respJson *UpdateRedirectUriResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type UpdateRedirectUriParams struct {
	BrandingId  int64
	Description string

	Uri string
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/deleteredirecturi/bd9ecc691a0f7a1c2281a4bc737bb756
func DeleteRedirectUri(appkey string, redirecturiid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/redirecturis/{RedirectUriId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/createredirecturi/427c540d153b3228eea8974e5bc0d792
func CreateRedirectUri(appkey string, params CreateRedirectUriParams) (*CreateRedirectUriResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}/redirecturis")
	if err != nil {
		return nil, err
	}
	var respJson *CreateRedirectUriResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CreateRedirectUriParams struct {
	BrandingId  int64
	Description string
	Uri         string // required
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getapps/e6de279d5c2f627b90786a21ccd41d65
func GetApps(params GetAppsParams) (*GetAppsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAppsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAppsParams struct {
	inlinecount string
	skip        int64
	top         int64
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getappbyappkey/e24dc3c2fbadadefa322dc329ac555c4
func GetAppByAppKey(appkey string) (*GetAppByAppKeyResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAppByAppKeyResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/getresourcebyappkey/7c4bfc00847280ed6096e6b6becfdd9b
func GetResourceByAppKey(appkey string) (*GetResourceByAppKeyResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/apps/resource/{AppKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetResourceByAppKeyResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/createapp/263cc3cc159896761d1de92b8fe051f2
func CreateApp(params CreateAppParams) (*CreateAppResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/developer/apps")
	if err != nil {
		return nil, err
	}
	var respJson *CreateAppResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CreateAppParams struct {
	Description      string // required
	Flow             string // required
	IsTradingEnabled bool
	Name             string // required
	RedirectUri      string
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/updateapp/f078614ace426838be44c3fafdeeeec2
func UpdateApp(appkey string, params UpdateAppParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateAppParams struct {
	Description string
	Name        string
}

// https://www.developer.saxo/openapi/referencedocs/developer/v1/apps/deactivateapp/ac1b081340bcc1ef6195260b01ea014f
func DeactivateApp(appkey string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/developer/apps/{AppKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
