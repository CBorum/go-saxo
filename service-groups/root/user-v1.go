package root

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/root/v1/user/userinfo/a5a7f20c716323d335335dee22b637e7
func UserInfo() (*UserInfoResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/root/v1/user")
	if err != nil {
		return nil, err
	}
	var respJson *UserInfoResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
