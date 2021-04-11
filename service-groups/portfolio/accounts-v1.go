package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/getaccount/f692a2e73dafc187584df63f5a4b56fa
func GetAccount(accountkey string) (*GetAccountResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/{AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAccountResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/getaccounts/af56e3512758f8125dc6e5493d93c019
func GetAccountsMe(params GetAccountsMeParams) (*GetAccountsMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/me/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAccountsMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAccountsMeParams struct {
	inlinecount string
	skip        int64
	top         int64
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/getaccounts/ad32cd71ff2773e74f1ff333ab122f3c
func GetAccounts(params GetAccountsParams) (*GetAccountsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}&ClientKey={ClientKey}&IncludeSubAccounts={IncludeSubAccounts}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAccountsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAccountsParams struct {
	inlinecount        string
	skip               int64
	top                int64
	ClientKey          string // required
	IncludeSubAccounts bool
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/updateaccount/0d85e8834628f886d3345c998fe876a3
func UpdateAccount(accountkey string, params UpdateAccountParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/{AccountKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateAccountParams struct {
	AccountValueProtectionLimit        float64
	BenchmarkInstrument                string
	DisplayName                        string
	UseCashPositionsAsMarginCollateral bool
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/resetaccount/3b6e0d7c0ba729645f75e4f217390b5c
func ResetAccount(accountkey string, params ResetAccountParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/{AccountKey}/reset")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type ResetAccountParams struct {
	NewBalance float64 // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/addactivesubscription/4d9791c2ef0ae03fc4dac7c3ecc6472d
func AddActiveSubscriptionAccounts(params AddActiveSubscriptionAccountsParams) (*AddActiveSubscriptionAccountsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddActiveSubscriptionAccountsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddActiveSubscriptionAccountsParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/deletesubscriptions/c34094012096719dc56a9a0d4e2594f3
func DeleteSubscriptionsAccounts(contextid string, params DeleteSubscriptionsAccountsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsAccountsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accounts/deletesubscription/5aabe67e8bd7901c4128ec2b9605b701
func DeleteSubscriptionAccounts(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/accounts/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
