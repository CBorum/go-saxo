package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetposition/cd1c10ce006e0c9f4db7c1f490191d51
func GetNetPosition(netpositionid string, params GetNetPositionParams) (*GetNetPositionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/{NetPositionId}/?FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetNetPositionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
	FieldGroups     string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetpositiondetails/fe956944f8bb97cc409525a0f2281f56
func GetNetPositionDetails(netpositionid string, params GetNetPositionDetailsParams) (*GetNetPositionDetailsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/{NetPositionId}/details/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetNetPositionDetailsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionDetailsParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required

}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetpositions/77f81ee9488691ec98679a830bb738c4
func GetNetPositionsMe(params GetNetPositionsMeParams) (*GetNetPositionsMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/me/?$top={$top}&$skip={$skip}&FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetNetPositionsMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionsMeParams struct {
	skip        int64
	top         int64
	FieldGroups string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/getnetpositions/372395024bc5f81341194857f088364e
func GetNetPositions(params GetNetPositionsParams) (*GetNetPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/?$top={$top}&$skip={$skip}&NetPositionId={NetPositionId}&Uic={Uic}&AssetType={AssetType}&ValueDate={ValueDate}&ExpiryDate={ExpiryDate}&PutCall={PutCall}&Strike={Strike}&UpperBarrier={UpperBarrier}&LowerBarrier={LowerBarrier}&WatchlistId={WatchlistId}&FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetNetPositionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetNetPositionsParams struct {
	skip            int64
	top             int64
	AccountGroupKey string
	AccountKey      string
	AssetType       string
	ClientKey       string // required
	ExpiryDate      string
	FieldGroups     string
	LowerBarrier    string
	NetPositionId   string
	PutCall         string
	Strike          string
	Uic             int64
	UpperBarrier    string
	ValueDate       string
	WatchlistId     string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/addactivesubscription/acbe15412d603d0b3971695cb0bb5312
func AddActiveSubscriptionNetPositions(params AddActiveSubscriptionNetPositionsParams) (*AddActiveSubscriptionNetPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddActiveSubscriptionNetPositionsResponse
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
func DeleteSubscriptionsNetPositions(contextid string, params DeleteSubscriptionsNetPositionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsNetPositionsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/netpositions/deletesubscription/f9915a776e787f4d05f8820d1cfd0cf0
func DeleteSubscriptionNetPositions(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/netpositions/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
