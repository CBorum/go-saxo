// This file is autogenerated by cmd/generate/main.go
package root

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/root/v1/subscriptions/deletesubscriptions/14d1d2d03f6aca45ab9c29e38973f449
func DeleteSubscriptions(contextid string, params *DeleteSubscriptionsParams) ([]byte, error) {
    url := "https://gateway.saxobank.com/sim/openapi/root/v1/subscriptions/{ContextId}/?Tag={Tag}"
    url = saxo.PrepareUrlRoute(url, saxo.RP("{ContextId}", contextid))
    url = saxo.PrepareUrlParams(url, params)
    resp, err := saxo.GetClient().DoRequest("DELETE", url, nil) 
    if err != nil {
        return nil, err
    }
    return resp.Bytes(), nil 
}

type DeleteSubscriptionsParams struct { 
    
    Tag string 
}
