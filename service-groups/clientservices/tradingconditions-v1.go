package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/tradingconditions/getforinstrument/f5c6f7adfb6d1230ea9cb95b801ed806
func GetForInstrument(accountkey string, assettype string, uic string, params GetForInstrumentParams) (*GetForInstrumentResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/tradingconditions/instrument/{AccountKey}/{Uic}/{AssetType}/?FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetForInstrumentResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetForInstrumentParams struct {
	FieldGroups string
}
