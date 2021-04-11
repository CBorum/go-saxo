package marketoverview

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/mkt/v1/instrumentdocument/getrecommendedinstrumentdocumentinformation/2c4d975085f24f29a7ccf2cf240c6af6
func GetRecommendedInstrumentDocumentInformationv1(assettype string, uic string, params GetRecommendedInstrumentDocumentInformationv1Params) (*GetRecommendedInstrumentDocumentInformationResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/mkt/v1/instruments/{Uic}/{AssetType}/documents/recommended/?DocumentType={DocumentType}")
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

type GetRecommendedInstrumentDocumentInformationv1Params struct {
	DocumentType string // required

}

// https://www.developer.saxo/openapi/referencedocs/mkt/v1/instrumentdocument/getdocumentbydocumenttypeandlanguagecode/bb48e9bd37dfafb53ae73d36807c43db
func GetDocumentByDocumentTypeAndLanguageCodev1(assettype string, uic string, params GetDocumentByDocumentTypeAndLanguageCodev1Params) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/mkt/v1/instruments/{Uic}/{AssetType}/documents/pdf/?DocumentType={DocumentType}&LanguageCode={LanguageCode}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GetDocumentByDocumentTypeAndLanguageCodev1Params struct {
	DocumentType string // required
	LanguageCode string // required

}
