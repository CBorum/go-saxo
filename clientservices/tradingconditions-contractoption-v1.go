// This file is autogenerated by cmd/generate/main.go
package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/tradingconditions-contractoption/getforcontractoption/6556258a209cfef88493b8c47a92257e
func GetForContractOption(accountkey string, optionrootid string, params *GetForContractOptionParams) (*GetForContractOptionResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/tradingconditions/ContractOptionSpaces/{AccountKey}/{OptionRootId}/?FieldGroups={FieldGroups}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{AccountKey}", accountkey), saxo.RP("{OptionRootId}", optionrootid))
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetForContractOptionResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetForContractOptionParams struct {
	FieldGroups string `url:",omitempty"`
}
