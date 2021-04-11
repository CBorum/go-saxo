package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/tradingconditions-contractoption/getforcontractoption/6556258a209cfef88493b8c47a92257e
func GetForContractOption(accountkey string, optionrootid string, params GetForContractOptionParams) (*GetForContractOptionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/tradingconditions/ContractOptionSpaces/{AccountKey}/{OptionRootId}/?FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetForContractOptionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetForContractOptionParams struct {
	FieldGroups string
}
