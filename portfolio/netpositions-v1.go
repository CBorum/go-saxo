// This file is autogenerated by cmd/generate/main.go
package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetposition/cd1c10ce006e0c9f4db7c1f490191d51
func GetNetPosition(netpositionid string, params *GetNetPositionParams) (*GetNetPositionResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/{NetPositionId}/?FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{NetPositionId}", netpositionid))
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetNetPositionResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionParams struct {
	AccountGroupKey string `url:",omitempty"`
	AccountKey      string `url:",omitempty"`
	ClientKey       string // required
	FieldGroups     string `url:",omitempty"`
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetpositiondetails/fe956944f8bb97cc409525a0f2281f56
func GetNetPositionDetails(netpositionid string, params *GetNetPositionDetailsParams) (*GetNetPositionDetailsResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/{NetPositionId}/details/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{NetPositionId}", netpositionid))
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetNetPositionDetailsResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionDetailsParams struct {
	AccountGroupKey string `url:",omitempty"`
	AccountKey      string `url:",omitempty"`
	ClientKey       string // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetpositions/77f81ee9488691ec98679a830bb738c4
func GetNetPositionsMe(params *GetNetPositionsMeParams) (*GetNetPositionsMeResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/me/?$top={$top}&$skip={$skip}&FieldGroups={FieldGroups}"
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetNetPositionsMeResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionsMeParams struct {
	Skip        int64  `url:",omitempty"`
	Top         int64  `url:",omitempty"`
	FieldGroups string `url:",omitempty"`
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetpositions/372395024bc5f81341194857f088364e
func GetNetPositions(params *GetNetPositionsParams) (*GetNetPositionsResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/?$top={$top}&$skip={$skip}&NetPositionId={NetPositionId}&Uic={Uic}&AssetType={AssetType}&ValueDate={ValueDate}&ExpiryDate={ExpiryDate}&PutCall={PutCall}&Strike={Strike}&UpperBarrier={UpperBarrier}&LowerBarrier={LowerBarrier}&WatchlistId={WatchlistId}&FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}"
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetNetPositionsResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionsParams struct {
	Skip            int64  `url:",omitempty"`
	Top             int64  `url:",omitempty"`
	AccountGroupKey string `url:",omitempty"`
	AccountKey      string `url:",omitempty"`
	AssetType       string `url:",omitempty"`
	ClientKey       string // required
	ExpiryDate      string `url:",omitempty"`
	FieldGroups     string `url:",omitempty"`
	LowerBarrier    string `url:",omitempty"`
	NetPositionId   string `url:",omitempty"`
	PutCall         string `url:",omitempty"`
	Strike          string `url:",omitempty"`
	Uic             int64  `url:",omitempty"`
	UpperBarrier    string `url:",omitempty"`
	ValueDate       string `url:",omitempty"`
	WatchlistId     string `url:",omitempty"`
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/addactivesubscription/acbe15412d603d0b3971695cb0bb5312
func AddActiveSubscriptionNetPositions(params *AddActiveSubscriptionNetPositionsParams) (*AddActiveSubscriptionNetPositionsResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/subscriptions"
	resp, err := saxo.GetClient().DoRequest("POST", url, params)
	if err != nil {
		return nil, err
	}
	respJson := &AddActiveSubscriptionNetPositionsResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddActiveSubscriptionNetPositionsParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/deletesubscriptions/fbe06830552db20330e85b1c97308970
func DeleteSubscriptionsNetPositions(contextid string, params *DeleteSubscriptionsNetPositionsParams) ([]byte, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/subscriptions/{ContextId}/?Tag={Tag}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{ContextId}", contextid))
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsNetPositionsParams struct {
	Tag string `url:",omitempty"`
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/deletesubscription/f9915a776e787f4d05f8820d1cfd0cf0
func DeleteSubscriptionNetPositions(contextid string, referenceid string) ([]byte, error) {
	url := "https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/subscriptions/{ContextId}/{ReferenceId}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{ContextId}", contextid), saxo.RP("{ReferenceId}", referenceid))
	resp, err := saxo.GetClient().DoRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
