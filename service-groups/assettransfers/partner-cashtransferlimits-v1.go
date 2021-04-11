package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/partner-cashtransferlimits/getcashtransferslimits/1fe73536a26078ae94af96a6fc33268d
func GetCashTransfersLimits(params GetCashTransfersLimitsParams) (*GetCashTransfersLimitsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/partner/cashtransferlimits/?ClientKey={ClientKey}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetCashTransfersLimitsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetCashTransfersLimitsParams struct {
	AccountKey string // required
	ClientKey  string // required
}
