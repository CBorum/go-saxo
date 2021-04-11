package root

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/root/v1/sessions/getcapabilities/7fe2af209d7d1d351faa2774702c749c
func GetCapabilities() (*GetCapabilitiesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/root/v1/sessions/capabilities")
	if err != nil {
		return nil, err
	}
	var respJson *GetCapabilitiesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/sessions/setcapabilities/0d4fbb3253171f8979462a11e91c1237
func SetCapabilities(params SetCapabilitiesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/root/v1/sessions/capabilities")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type SetCapabilitiesParams struct {
	TradeLevel string
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/sessions/patchcapabilities/0406615b3eb2adb17fa24f870f0b9202
func PatchCapabilities(params PatchCapabilitiesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/root/v1/sessions/capabilities")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type PatchCapabilitiesParams struct {
	TradeLevel string
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/sessions/addsubscription/319e322e524340b7caa662adac83541a
func AddSubscription(params AddSubscriptionParams) (*AddSubscriptionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/root/v1/sessions/events/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionParams struct {
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/sessions/deletesubscription/a7d07e55e77d4bfd3e52d703595ab76c
func DeleteSubscription(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/root/v1/sessions/events/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
