package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/cashmanagement-beneficiaryinstructions/getbeneficiaryinstructionsasync/d1fada87658f83fb9ea6772312cf4b13
func GetBeneficiaryInstructionsAsync(params GetBeneficiaryInstructionsAsyncParams) (*GetBeneficiaryInstructionsAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/cashmanagement/beneficiaryinstructions/?ClientKey={ClientKey}&Currency={Currency}")
	if err != nil {
		return nil, err
	}
	var respJson *GetBeneficiaryInstructionsAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetBeneficiaryInstructionsAsyncParams struct {
	ClientKey string
	Currency  string
}
