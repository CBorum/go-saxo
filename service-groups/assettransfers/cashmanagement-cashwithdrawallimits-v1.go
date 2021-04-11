package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/cashmanagement-cashwithdrawallimits/getcashwithdrawallimits/864282f824520df83614e7d4287c0979
func GetCashWithdrawalLimits(params GetCashWithdrawalLimitsParams) (*GetCashWithdrawalLimitsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/cashmanagement/withdrawallimits/?ClientKey={ClientKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetCashWithdrawalLimitsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetCashWithdrawalLimitsParams struct {
	AccountKey string // required
	ClientKey  string // required
}
