package portfolio

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/getinstrumentsexposures/ab93eea87c584040f55510b93eeb75f2
func GetInstrumentsExposures() (*GetInstrumentsExposuresResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/instruments/me")
	if err != nil {
		return nil, err
	}
	var respJson *GetInstrumentsExposuresResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/getinstrumentsexposuresforloggedinuser/66dbe48120b9e4510567a7f44b7a8580
func GetInstrumentsExposuresForLoggedInUser(params GetInstrumentsExposuresForLoggedInUserParams) (*GetInstrumentsExposuresForLoggedInUserResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/instruments/?Uic={Uic}&AssetType={AssetType}&ValueDate={ValueDate}&ExpiryDate={ExpiryDate}&PutCall={PutCall}&Strike={Strike}&UpperBarrier={UpperBarrier}&LowerBarrier={LowerBarrier}&ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetInstrumentsExposuresForLoggedInUserResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetInstrumentsExposuresForLoggedInUserParams struct {
	AccountGroupKey string
	AccountKey      string
	AssetType       string
	ClientKey       string // required
	ExpiryDate      string
	LowerBarrier    string
	PutCall         string
	Strike          string
	Uic             int64
	UpperBarrier    string
	ValueDate       string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/addactivesubscription/d6963365a060776b79d3e9713ef0a7e2
func AddActiveSubscriptionExposures(params AddActiveSubscriptionExposuresParams) (*AddActiveSubscriptionExposuresResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/instruments/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddActiveSubscriptionExposuresResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddActiveSubscriptionExposuresParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/deletesubscriptions/0f53c7d5b8c3a22c4f34bf6034a8de61
func DeleteSubscriptionsExposures(contextid string, params DeleteSubscriptionsExposuresParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/instruments/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsExposuresParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/deletesubscription/c80bfdfa57b45923050bfc437bc52ac9
func DeleteSubscriptionExposures(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/instruments/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/getcurrencyexposures/bd3b48164105531dcba0069db33a27ec
func GetCurrencyExposures() (*GetCurrencyExposuresResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/currency/me")
	if err != nil {
		return nil, err
	}
	var respJson *GetCurrencyExposuresResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/getcurrencyexposuresforloggedinuser/3ef64ff498143552f7335457ba98c4c9
func GetCurrencyExposuresForLoggedInUser(params GetCurrencyExposuresForLoggedInUserParams) (*GetCurrencyExposuresForLoggedInUserResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/currency/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetCurrencyExposuresForLoggedInUserResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetCurrencyExposuresForLoggedInUserParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/getfxspotnetexposures/e97cc469e4ce12a66447c5a8d485181d
func GetFxSpotNetExposures() (*GetFxSpotNetExposuresResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/fxspot/me")
	if err != nil {
		return nil, err
	}
	var respJson *GetFxSpotNetExposuresResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/port/v1/exposure/getfxspotnetexposuresforloggedinuser/48a55afbadbf96ede73ef5f99e29c803
func GetFxSpotNetExposuresForLoggedInUser(params GetFxSpotNetExposuresForLoggedInUserParams) (*GetFxSpotNetExposuresForLoggedInUserResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/port/v1/exposure/fxspot/?ClientKey={ClientKey}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetFxSpotNetExposuresForLoggedInUserResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetFxSpotNetExposuresForLoggedInUserParams struct {
	AccountGroupKey string
	AccountKey      string
	ClientKey       string // required
}
