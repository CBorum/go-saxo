package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/partner-cashtransfer/cashtransfer/cddc7ce8c4e6757f9f64566a46ccdbb1
func CashTransfer(params CashTransferParams) (*CashTransferResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/atr/v1/partner/cashtransfers")
	if err != nil {
		return nil, err
	}
	var respJson *CashTransferResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CashTransferParams struct {
	Amount            float64 // required
	Comment           string
	Currency          string // required
	ExternalReference string
	FromAccountKey    string // required
	FundingCheck      string
	ToAccountKey      string // required
	ValueDate         string
}

// https://www.developer.saxo/openapi/referencedocs/atr/v1/partner-cashtransfer/cashtransfersbyid/5de341702cf93764cccd447c07d4f55a
func CashTransfersById(transactionid string) (*CashTransfersByIdResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/partner/cashtransfers/{TransactionId}")
	if err != nil {
		return nil, err
	}
	var respJson *CashTransfersByIdResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/atr/v1/partner-cashtransfer/cashtransfersbysearchparameters/49eb5f4687599a6494ab2cd51edeb7d2
func CashTransfersBySearchParameters(params CashTransfersBySearchParametersParams) (*CashTransfersBySearchParametersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/atr/v1/partner/cashtransfers/?AccountKey={AccountKey}&ClientKey={ClientKey}&FromDateTime={FromDateTime}&ToDateTime={ToDateTime}&ExternalReference={ExternalReference}&Top={Top}&SkipToken={SkipToken}")
	if err != nil {
		return nil, err
	}
	var respJson *CashTransfersBySearchParametersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CashTransfersBySearchParametersParams struct {
	AccountKey        string
	ClientKey         string
	ExternalReference string
	FromDateTime      string
	SkipToken         string
	ToDateTime        string
	Top               int64
}
