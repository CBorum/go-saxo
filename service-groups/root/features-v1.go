package root

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/root/v1/features/get/fa6ff69b66f9071d69d9e5088f26236c
func GetFeatures() (*GetFeaturesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/root/v1/features/availability")
	if err != nil {
		return nil, err
	}
	var respJson *GetFeaturesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/features/post/0f4a768cafca2fd5073757fade0a35d6
func PostFeatures(params PostFeaturesParams) (*PostFeaturesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/root/v1/features/availability/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *PostFeaturesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type PostFeaturesParams struct {
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/features/delete/91bcb26e5d95cf6ef1738f2736d2fb29
func DeleteFeatures(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/root/v1/features/availability/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
