// This file is autogenerated by cmd/generate/main.go
package clientreporting

import (
    "fmt"

	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cr/v1/historicalreportdata-portfoliomanagement/getasync/562203fc923449262bf0624e9284d66b
func GetAsyncPortfolioManagement(clientkey string, fromdate string, todate string, params *GetAsyncPortfolioManagementParams) ([]byte, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/cr/v1/reports/Portfolio/{ClientKey}/{Fromdate}/{ToDate}/?AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    return resp.Bytes(), nil 
}

type GetAsyncPortfolioManagementParams struct { 
    AccountGroupKey string
    AccountKey string
    
    
     
}

// https://www.developer.saxo/openapi/referencedocs/cr/v1/historicalreportdata-portfoliomanagement/getmeasync/466315d85a16246a7ab0f3e3c2ff516c
func GetMeAsyncPortfolioManagement(fromdate string, todate string, params *GetMeAsyncPortfolioManagementParams) ([]byte, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/cr/v1/reports/Portfolio/me/{Fromdate}/{ToDate}/?AccountGroupKey={AccountGroupKey}&AccountKey={AccountKey}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    return resp.Bytes(), nil 
}

type GetMeAsyncPortfolioManagementParams struct { 
    AccountGroupKey string
    AccountKey string
    
     
}
