package corporateactions

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ca/v1/proxyvoting/getproxyvotingeventsasync/fa0a55afb94926ed3fc7d756c07f8fa8
func GetProxyVotingEventsAsync(params GetProxyVotingEventsAsyncParams) (*GetProxyVotingEventsAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ca/v1/proxyvoting/events/?$top={$top}&$skip={$skip}&ClientKey={ClientKey}&SortType={SortType}&SortColumn={SortColumn}")
	if err != nil {
		return nil, err
	}
	var respJson *GetProxyVotingEventsAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetProxyVotingEventsAsyncParams struct {
	skip       int64
	top        int64
	ClientKey  string // required
	SortColumn string
	SortType   string
}
