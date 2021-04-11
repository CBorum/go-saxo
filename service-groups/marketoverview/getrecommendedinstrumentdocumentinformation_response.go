package marketoverview

type GetRecommendedInstrumentDocumentInformationResponse struct {
	Data []struct {
		DocumentDateTime   string  `json:"DocumentDateTime"`
		DocumentRelationID float64 `json:"DocumentRelationId"`
		DocumentType       string  `json:"DocumentType"`
		LanguageCode       string  `json:"LanguageCode"`
	} `json:"Data"`
}
