package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v2/clientinfo/search/751324b669928ae0a4dbc8c28c003d03
func Search(params SearchParams) (*SearchResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cs/v2/clientinfo/clients/search/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}")
	if err != nil {
		return nil, err
	}
	var respJson *SearchResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type SearchParams struct {
	inlinecount string
	skip        int64
	top         int64
	AccountId   string
	AccountKey  string
	ClientId    string
	ClientKey   string
	FieldGroups string
	Keywords    string
	UserId      string
}
