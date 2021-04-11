package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/getposition/2834dbe214c4a4a8027a55a62ce0fa52
func GetPosition(positionid string, params GetPositionParams) (*GetPositionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/positions/{PositionId}/?FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetPositionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetPositionParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
	FieldGroups     string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/getpositiondetails/99b850939cbc7b0fa762221b27df1bf7
func GetPositionDetails(positionid string, params GetPositionDetailsParams) (*GetPositionDetailsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/positions/{PositionId}/details/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetPositionDetailsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetPositionDetailsParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required

}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/getpositions/b6d549a50a5f35244806aaa0554d6eae
func GetPositionsMe(params GetPositionsMeParams) (*GetPositionsMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/positions/me/?$top={$top}&$skip={$skip}&FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetPositionsMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetPositionsMeParams struct {
	skip        int64
	top         int64
	FieldGroups string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/getpositions/bfbf88ed3670fed36d07a558cf317254
func GetPositions(params GetPositionsParams) (*GetPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/positions/?$top={$top}&$skip={$skip}&NetPositionId={NetPositionId}&PositionId={PositionId}&WatchlistId={WatchlistId}&FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetPositionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetPositionsParams struct {
	skip            int64
	top             int64
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
	FieldGroups     string
	NetPositionId   string
	PositionId      string
	WatchlistId     string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/addactivesubscription/ccdff431c0efe7f9392a80239f7955d6
func AddActiveSubscriptionPositions(params AddActiveSubscriptionPositionsParams) (*AddActiveSubscriptionPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/positions/subscriptions/?$top={$top}")
	if err != nil {
		return nil, err
	}
	var respJson *AddActiveSubscriptionPositionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddActiveSubscriptionPositionsParams struct {
	top         int64
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/updatepositionsubscriptionpagesize/ba58baaa810ca77c20fd72b44df50821
func UpdatePositionSubscriptionPagesize(contextid string, referenceid string, params UpdatePositionSubscriptionPagesizeParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/positions/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdatePositionSubscriptionPagesizeParams struct {
	NewPageSize int64
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/deletesubscriptions/cec29ec11e88b1d530b675fd395f916c
func DeleteSubscriptionsPositions(contextid string, params DeleteSubscriptionsPositionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/positions/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsPositionsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/positions/deletesubscription/cd8cd11da69274bbc401e46238fb287e
func DeleteSubscriptionPositions(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/positions/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
