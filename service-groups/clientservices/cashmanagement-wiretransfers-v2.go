package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v2/cashmanagement-wiretransfers/get/d101bb71d038dfb9487f4ede1fa0186b
func GetWireTransfers(params GetWireTransfersParams) (*GetWireTransfersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v2/cashmanagement/wiretransfers/instructions/?ClientKey={ClientKey}&AccountKey={AccountKey}&CurrencyCode={CurrencyCode}")
	if err != nil {
		return nil, err
	}
	var respJson *GetWireTransfersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetWireTransfersParams struct {
	AccountKey   string // required
	ClientKey    string // required
	CurrencyCode string // required
}
