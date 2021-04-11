package marketoverview

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/mkt/v2/instrumentdocument/getrecommendedinstrumentdocumentinformation/7663849659e774f55e6a9005c0ff7832
func GetRecommendedInstrumentDocumentInformationv2(assettype string, uic string, params GetRecommendedInstrumentDocumentInformationv2Params) (*GetRecommendedInstrumentDocumentInformationResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/mkt/v2/instruments/{Uic}/{AssetType}/documents/recommended/?DocumentTypes={DocumentTypes}")
	if err != nil {
		return nil, err
	}
	var respJson *GetRecommendedInstrumentDocumentInformationResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetRecommendedInstrumentDocumentInformationv2Params struct {
	DocumentTypes string // required

}

// https://www.developer.saxo/openapi/referencedocs/mkt/v2/instrumentdocument/getdocumentbydocumenttypeandlanguagecode/28a08aeaa913ed3e88911e073c43ccb1
func GetDocumentByDocumentTypeAndLanguageCodev2(assettype string, uic string, params GetDocumentByDocumentTypeAndLanguageCodev2Params) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/mkt/v2/instruments/{Uic}/{AssetType}/documents/pdf/?DocumentType={DocumentType}&LanguageCode={LanguageCode}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GetDocumentByDocumentTypeAndLanguageCodev2Params struct {
	DocumentType string // required
	LanguageCode string // required

}
