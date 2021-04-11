package clientmanagement

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cm/v1/users/resetpassword/9f8c2cf85499f04f15adf43e76bf073a
func ResetPassword(params ResetPasswordParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cm/v1/users/resetpasswordrequest")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type ResetPasswordParams struct {
	Email    string // required
	Language string
	UserId   string // required
}
