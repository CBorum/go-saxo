package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/securitiestransfers/transfer/cf6e251363afbbbe5660acfb9c847bc6
func Transfer(params TransferParams) (*TransferResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/atr/v1/securitiestransfers/transfers")
	if err != nil {
		return nil, err
	}
	var respJson *TransferResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type TransferParams struct {
	AccountKey            string // required
	AssetType             string // required
	Broker                string // required
	ClientAccountAtBroker string // required
	ClientKey             string // required
	Securities            string // required
	SettlementDate        string
	SettlementInstruction string
	TradeDate             string
	TransferDirection     string // required
}

// https://www.developer.saxo/openapi/referencedocs/atr/v1/securitiestransfers/getbrokers/22b7d795c89eb4fa20b06f525a6d58a4
func GetBrokers(params GetBrokersParams) (*GetBrokersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/securitiestransfers/brokers/?CountryCode={CountryCode}")
	if err != nil {
		return nil, err
	}
	var respJson *GetBrokersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetBrokersParams struct {
	CountryCode string // required
}

// https://www.developer.saxo/openapi/referencedocs/atr/v1/securitiestransfers/gettransferdetails/5961f7c694770ea83d4190909e1c4c89
func GetTransferDetails(params GetTransferDetailsParams) (*GetTransferDetailsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/securitiestransfers/transfers/?TransferIds={TransferIds}&ClientKey={ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTransferDetailsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetTransferDetailsParams struct {
	ClientKey   string
	TransferIds string // required
}
