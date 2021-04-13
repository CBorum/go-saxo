package portfolio

type GetFxSpotNetExposuresResponse struct {
	Data []struct {
		Amount                            float64 `json:"Amount"`
		AmountInCalculationEntityCurrency float64 `json:"AmountInCalculationEntityCurrency"`
		Currency                          string  `json:"Currency"`
	} `json:"Data"`
	MaxRows float64 `json:"MaxRows"`
}
