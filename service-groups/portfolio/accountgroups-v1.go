package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/accountgroups/getaccountgroup/052ea6c89788fa97642457e08d63ba7f
func GetAccountGroup(accountgroupkey string, params GetAccountGroupParams) (*GetAccountGroupResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/accountgroups/{AccountGroupKey}/?ClientKey={ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAccountGroupResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAccountGroupParams struct {
	ClientKey string // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accountgroups/getaccountgroupsforloggedinuser/9ce455f6be24a5398da000b453611214
func GetAccountGroupsForLoggedInUser(params GetAccountGroupsForLoggedInUserParams) (*GetAccountGroupsForLoggedInUserResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/accountgroups/me/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAccountGroupsForLoggedInUserResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAccountGroupsForLoggedInUserParams struct {
	inlinecount string
	skip        int64
	top         int64
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accountgroups/getaccountgroups/c597d192d53a83c7090d806242eb4902
func GetAccountGroups(params GetAccountGroupsParams) (*GetAccountGroupsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/accountgroups/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}&ClientKey={ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAccountGroupsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAccountGroupsParams struct {
	inlinecount string
	skip        int64
	top         int64
	ClientKey   string // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/accountgroups/updateaccountgroup/9ffe50fd5e042cd8f557983d370bd3ac
func UpdateAccountGroup(accountgroupkey string, params UpdateAccountGroupParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/accountgroups/{AccountGroupKey}/?ClientKey={ClientKey}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateAccountGroupParams struct {
	AccountValueProtectionLimit float64
	ClientKey                   string // required
}
