package clientmanagement

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cm/v2/signups/attachfile/c454ca859a123e91d13713077160e4bf
func AttachFilev2(signupid string, params AttachFilev2Params) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cm/v2/signups/attachments/{SignUpId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type AttachFilev2Params struct {
	Documents string
}

// https://www.developer.saxo/openapi/referencedocs/cm/v2/signups/getsignupoptions/ca505c6067de94f36e04ab9c37895413
func GetSignupOptionsv2() (*GetSignupOptionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cm/v2/signups/options")
	if err != nil {
		return nil, err
	}
	var respJson *GetSignupOptionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
