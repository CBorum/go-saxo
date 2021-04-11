package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/tradingconditions-cost/gettradingconditioncost/665e2b2d64ea678fbf102d40fd82a3ec
func GetTradingConditionCost(accountkey string, assettype string, uic string, params GetTradingConditionCostParams) (*GetTradingConditionCostResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/tradingconditions/cost/{AccountKey}/{Uic}/{AssetType}/?Amount={Amount}&FieldGroups={FieldGroups}&Price={Price}&HoldingPeriodInDays={HoldingPeriodInDays}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradingConditionCostResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetTradingConditionCostParams struct {
	Amount float64 // required

	FieldGroups         string
	HoldingPeriodInDays int64
	Price               float64
}
