package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/users/getuser/24acd0f6657d7f2a34f4a37ccad185f7
func GetUserMe() (*GetUserMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/users/me")
	if err != nil {
		return nil, err
	}
	var respJson *GetUserMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/users/getusers/6b4194594e4b90d95179a448eede0cbd
func GetUsers(params GetUsersParams) (*GetUsersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/users/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}&ClientKey={ClientKey}&IncludeSubUsers={IncludeSubUsers}")
	if err != nil {
		return nil, err
	}
	var respJson *GetUsersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetUsersParams struct {
	inlinecount     string
	skip            int64
	top             int64
	ClientKey       string // required
	IncludeSubUsers bool
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/users/getuser/254eac9c10b25ad7b042379fd2c60b02
func GetUser(userkey string) (*GetUserResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/users/{UserKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetUserResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/users/updateuserpreferences/a6d4ba6f6359975cdfa2195b28554d73
func UpdateUserPreferences(params UpdateUserPreferencesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/users/me")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateUserPreferencesParams struct {
	Culture    string
	Language   string
	TimeZoneId int64
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/users/getallcliententitlements/c68178c4aeb8672eae6e08b8b5cc92e6
func GetAllClientEntitlements() (*GetAllClientEntitlementsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/users/me/entitlements")
	if err != nil {
		return nil, err
	}
	var respJson *GetAllClientEntitlementsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
