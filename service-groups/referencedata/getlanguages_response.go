package referencedata

type GetLanguagesResponse struct {
	Data []struct {
		LanguageCode string `json:"LanguageCode"`
		LanguageName string `json:"LanguageName"`
		NativeName   string `json:"NativeName"`
	} `json:"Data"`
	Next string `json:"__next"`
}
