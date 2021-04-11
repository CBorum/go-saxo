package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/audit-activities/get/65af3f977d65b3d98c66d7e86aa3ef8e
func GetActivities(params GetActivitiesParams) (*GetActivitiesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cs/v1/audit/activities/?$top={$top}&$skiptoken={$skiptoken}&ClientKey={ClientKey}&From={From}&To={To}&LogEntryTypes={LogEntryTypes}&AccountKey={AccountKey}&PostingId={PostingId}&CorrelationKey={CorrelationKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetActivitiesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetActivitiesParams struct {
	skiptoken      string
	top            int64
	AccountKey     string
	ClientKey      string // required
	CorrelationKey string
	From           string
	LogEntryTypes  string
	PostingId      string
	To             string
}
