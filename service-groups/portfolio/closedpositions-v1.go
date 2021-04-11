package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/getclosedpositions/bedb618e2ee9b8f8e88209b57493e939
func GetClosedPositions(params GetClosedPositionsParams) (*GetClosedPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/?$top={$top}&$skip={$skip}&ClosedPositionId={ClosedPositionId}&FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClosedPositionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClosedPositionsParams struct {
	skip             int64
	top              int64
	AccountGroupKey  string
	AccountKey       string
	ClientKey        string // required
	ClosedPositionId string
	FieldGroups      string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/getclosedposition/d3c1eab637010860015450533fc94282
func GetClosedPosition(closedpositionid string, params GetClosedPositionParams) (*GetClosedPositionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/{ClosedPositionId}/?FieldGroups={FieldGroups}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClosedPositionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClosedPositionParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required

	FieldGroups string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/getpositiondetails/0c2aeff81184a3db67698612160a72ac
func GetClosedPositionDetails(closedpositionid string, params GetClosedPositionDetailsParams) (*GetClosedPositionDetailsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/{ClosedPositionId}/details/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClosedPositionDetailsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClosedPositionDetailsParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required

}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/getclosedpositions/cec2dfd90666dd31517eee39e9fd15f8
func GetClosedPositionsMe(params GetClosedPositionsMeParams) (*GetClosedPositionsMeResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/me/?$top={$top}&$skip={$skip}&FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetClosedPositionsMeResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetClosedPositionsMeParams struct {
	skip        int64
	top         int64
	FieldGroups string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/addactivesubscription/39b4b705654a5f688c671badd6ec0a4c
func AddActiveSubscriptionClosedPositions(params AddActiveSubscriptionClosedPositionsParams) (*AddActiveSubscriptionClosedPositionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/subscriptions/?$top={$top}")
	if err != nil {
		return nil, err
	}
	var respJson *AddActiveSubscriptionClosedPositionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddActiveSubscriptionClosedPositionsParams struct {
	top         int64
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/updateclosedpositionsubscriptionpagesize/79a01dfff1294c6bc8a212e582b95b66
func UpdateClosedPositionSubscriptionPagesize(contextid string, referenceid string, params UpdateClosedPositionSubscriptionPagesizeParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateClosedPositionSubscriptionPagesizeParams struct {
	NewPageSize int64
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/deletesubscriptions/a713451058ef91dfb518817e87973dae
func DeleteSubscriptionsClosedPositions(contextid string, params DeleteSubscriptionsClosedPositionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsClosedPositionsParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/closedpositions/deletesubscription/ec42cb58a691ed57968b966e1d6b1f1b
func DeleteSubscriptionClosedPositions(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/closedpositions/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}