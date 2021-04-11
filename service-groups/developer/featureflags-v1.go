package developer

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/developer/v1/featureflags/getflags/e98a366b4bf75fac7af4bb9109357af1
func GetFlags() (*GetFlagsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/developer/featureflags")
	if err != nil {
		return nil, err
	}
	var respJson *GetFlagsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
