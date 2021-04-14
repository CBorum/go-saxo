// This file is autogenerated by cmd/generate/main.go
package valueadd

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/getallalertdefinitionsfilteredbystate/253503f194504d7e4192a9871c1ccef5
func GetAllAlertDefinitionsFilteredByState(params *GetAllAlertDefinitionsFilteredByStateParams) (*GetAllAlertDefinitionsFilteredByStateResponse, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/definitions/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}&State={State}"
    url = saxo.PrepareUrlParams(url, params)
    resp, err := saxo.GetClient().DoRequest("GET", url, nil) 
    if err != nil {
        return nil, err
    }
    respJson := &GetAllAlertDefinitionsFilteredByStateResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type GetAllAlertDefinitionsFilteredByStateParams struct { 
    inlinecount string
    skip int64
    top int64
    State string 
}

// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/getalertdefinition/2c8074078cc31b722a030eb5dcc36acc
func GetAlertDefinition(alertdefinitionid string) (*GetAlertDefinitionResponse, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/definitions/{AlertDefinitionId}"
    url = saxo.PrepareUrlRoute(url, saxo.RP("{AlertDefinitionId}", alertdefinitionid))
    resp, err := saxo.GetClient().DoRequest("GET", url, nil) 
    if err != nil {
        return nil, err
    }
    respJson := &GetAlertDefinitionResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}


// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/postalertdefinition/142c917d7073729471c019c8b14a032e
func PostAlertDefinition(params *PostAlertDefinitionParams) (*PostAlertDefinitionResponse, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/definitions"
    resp, err := saxo.GetClient().DoRequest("POST", url, nil) 
    if err != nil {
        return nil, err
    }
    respJson := &PostAlertDefinitionResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type PostAlertDefinitionParams struct { 
    AccountId string // required
    AssetType string // required
    Comment string
    ExpiryDate string // required
    IsRecurring bool // required
    Operator string // required
    PriceVariable string // required
    State string
    TargetValue float64 // required
    Uic int64 // required 
}

// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/putalertdefinition/5e45019d26e97ad5e7e5ec17f107982b
func PutAlertDefinition(alertdefinitionid string, params *PutAlertDefinitionParams) ([]byte, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/definitions/{AlertDefinitionId}"
    url = saxo.PrepareUrlRoute(url, saxo.RP("{AlertDefinitionId}", alertdefinitionid))
    resp, err := saxo.GetClient().DoRequest("PUT", url, nil) 
    if err != nil {
        return nil, err
    }
    return resp.Bytes(), nil 
}

type PutAlertDefinitionParams struct { 
    AccountId string // required
    
    AssetType string // required
    Comment string
    ExpiryDate string // required
    IsRecurring bool // required
    Operator string // required
    PriceVariable string // required
    State string
    TargetValue float64 // required
    Uic int64 // required 
}

// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/deletealertdefinitions/a50b7d0acffbb7c9a097a699b2749f98
func DeleteAlertDefinitions(alertdefinitionids string) ([]byte, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/definitions/{AlertDefinitionIds}"
    url = saxo.PrepareUrlRoute(url, saxo.RP("{AlertDefinitionIds}", alertdefinitionids))
    resp, err := saxo.GetClient().DoRequest("DELETE", url, nil) 
    if err != nil {
        return nil, err
    }
    return resp.Bytes(), nil 
}


// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/getusersettings/d7921c39174807f5c5b8dd52c8423a30
func GetUserSettings() (*GetUserSettingsResponse, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/usersettings"
    resp, err := saxo.GetClient().DoRequest("GET", url, nil) 
    if err != nil {
        return nil, err
    }
    respJson := &GetUserSettingsResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}


// https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/putusersettings/5f7047f70b89de00498121d47eae614b
func PutUserSettings(params *PutUserSettingsParams) ([]byte, error) {
    url := "https://gateway.saxobank.com/sim/openapi/vas/v1/pricealerts/usersettings"
    resp, err := saxo.GetClient().DoRequest("PUT", url, nil) 
    if err != nil {
        return nil, err
    }
    return resp.Bytes(), nil 
}

type PutUserSettingsParams struct { 
    EmailAddress string
    NotifyWithMail bool
    NotifyWithPopup bool
    Sound string 
}
