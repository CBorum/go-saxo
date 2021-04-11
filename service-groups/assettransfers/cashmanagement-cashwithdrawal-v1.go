package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/cashmanagement-cashwithdrawal/withdrawl/18f26c2a252695ffd0fb62314210bb8c
func Withdrawl(params WithdrawlParams) (*WithdrawlResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/atr/v1/cashmanagement/withdrawals")
	if err != nil {
		return nil, err
	}
	var respJson *WithdrawlResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type WithdrawlParams struct {
	AccountKey               string  // required
	Amount                   float64 // required
	BeneficiaryInstructionId string  // required
	Currency                 string  // required
	MessageToBeneficiary     string
}
