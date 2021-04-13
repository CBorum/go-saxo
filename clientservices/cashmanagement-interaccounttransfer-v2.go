// This file is autogenerated by cmd/generate/main.go
package clientservices

import (
    "fmt"

	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v2/cashmanagement-interaccounttransfer/transfer/da3c973663fa7a87842beab73a94c99d
func Transfer(params *TransferParams) (*TransferResponse, error) {
    resp, err := saxo.GetClient().DoRequest("POST", "https://gateway.saxobank.com/sim/openapi/cs/v2/cashmanagement/interaccounttransfers", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &TransferResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type TransferParams struct { 
    Amount float64 // required
    Currency string // required
    FromAccountKey string // required
    ToAccountKey string // required
    WithdrawalReasonId int64 
}