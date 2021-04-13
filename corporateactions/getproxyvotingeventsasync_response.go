package corporateactions

type GetProxyVotingEventsAsyncResponse struct {
	Data []struct {
		AccountID    string `json:"AccountId"`
		ActionURL    string `json:"ActionUrl"`
		Cins         string `json:"Cins"`
		Cusip        string `json:"Cusip"`
		CutoffDate   string `json:"CutoffDate"`
		DeliveryType struct {
			Code        string `json:"Code"`
			Description string `json:"Description"`
		} `json:"DeliveryType"`
		IsinCode   string `json:"IsinCode"`
		IssuerName string `json:"IssuerName"`
		JobNumber  string `json:"JobNumber"`
		Materials  []struct {
			MaterialType struct {
				Code        string `json:"Code"`
				Description string `json:"Description"`
			} `json:"MaterialType"`
			URL string `json:"Url"`
		} `json:"Materials"`
		ReceivedDate string `json:"ReceivedDate"`
		Status       struct {
			Code        string `json:"Code"`
			Description string `json:"Description"`
		} `json:"Status"`
		StatusDate string `json:"StatusDate"`
		Subtype    struct {
			Code        string `json:"Code"`
			Description string `json:"Description"`
		} `json:"Subtype"`
		Type struct {
			Code        string `json:"Code"`
			Description string `json:"Description"`
		} `json:"Type"`
	} `json:"Data"`
	Count float64 `json:"__count"`
}
