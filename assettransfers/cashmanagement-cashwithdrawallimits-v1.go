// This file is autogenerated by cmd/generate/main.go
package assettransfers

import (
    "fmt"

	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/cashmanagement-cashwithdrawallimits/getcashwithdrawallimits/864282f824520df83614e7d4287c0979
func GetCashWithdrawalLimits(params *GetCashWithdrawalLimitsParams) (*GetCashWithdrawalLimitsResponse, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/atr/v1/cashmanagement/withdrawallimits/?ClientKey={ClientKey}&AccountKey={AccountKey}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &GetCashWithdrawalLimitsResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type GetCashWithdrawalLimitsParams struct { 
    AccountKey string // required
    ClientKey string // required 
}