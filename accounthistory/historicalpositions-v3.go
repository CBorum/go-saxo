// This file is autogenerated by cmd/generate/main.go
package accounthistory

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/hist/v3/historicalpositions/gethistoricalpositions/390ff9543562ec3f62644714ba984c47
func GetHistoricalPositions(clientkey string, params *GetHistoricalPositionsParams) (*GetHistoricalPositionsResponse, error) {
    url := "https://gateway.saxobank.com/sim/openapi/hist/v3/positions/{ClientKey}/?$top={$top}&$skip={$skip}&$inlinecount={$inlinecount}&$skiptoken={$skiptoken}&Symbol={Symbol}&AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}&FromDate={FromDate}&ToDate={ToDate}&StandardPeriod={StandardPeriod}&AssetType={AssetType}&MockDataId={MockDataId}"
    url = saxo.PrepareUrlRoute(url, saxo.RP("{ClientKey}", clientkey))
    url = saxo.PrepareUrlParams(url, params)
    resp, err := saxo.GetClient().DoRequest("GET", url, nil) 
    if err != nil {
        return nil, err
    }
    respJson := &GetHistoricalPositionsResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type GetHistoricalPositionsParams struct { 
    inlinecount string
    skip int64
    skiptoken string
    top int64
    AccountGroupKey string
    AccountKey string
    AssetType string
    
    FromDate string
    MockDataId string
    StandardPeriod string
    Symbol string
    ToDate string 
}
