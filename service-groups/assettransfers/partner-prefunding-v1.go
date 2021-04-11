package assettransfers

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/atr/v1/partner-prefunding/prebookfunddeposit/f87719a93264749430d6eb7a234aa3af
func PrebookFundDeposit(params PrebookFundDepositParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/atr/v1/partner/prebookedfunds")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type PrebookFundDepositParams struct {
	AccountKey        string  // required
	Amount            float64 // required
	BIC               string  // required
	ClientKey         string  // required
	ClientName        string  // required
	Currency          string  // required
	ExpectedValueDate string  // required
	ExternalReference string  // required
	Iban              string  // required
	RemitterAccount   string  // required
}
