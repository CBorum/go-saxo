package trading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v1/optionschain/addsubscriptionasync/b6c2529e483362a3931cf8d76ca1ed15
func AddSubscriptionAsyncOptionsChain(params AddSubscriptionAsyncOptionsChainParams) (*AddSubscriptionAsyncOptionsChainResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/optionschain/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionAsyncOptionsChainResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionAsyncOptionsChainParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/optionschain/modifysubscription/708b5f9d261312d07249ef8b0c0ccf52
func ModifySubscriptionOptionsChain(contextid string, referenceid string, params ModifySubscriptionOptionsChainParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/trade/v1/optionschain/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type ModifySubscriptionOptionsChainParams struct {
	Expiries            string
	MaxStrikesPerExpiry int64
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/optionschain/deletesubscription/e1517f90a09669c3a5d1cc155a4985ca
func DeleteSubscriptionOptionsChain(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/optionschain/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/optionschain/resetsubscriptionatthemoney/7e907875f27f8443d73c8a6deb8efe31
func ResetSubscriptionAtTheMoney(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/trade/v1/optionschain/subscriptions/{ContextId}/{ReferenceId}/ResetATM")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
