package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v2/cashmanagement-cashwithdrawal/withdrawibanswift/2e34612be404d9a96b3de393e3d2ccf0
func WithdrawIBanSwift(params WithdrawIBanSwiftParams) (*WithdrawIBanSwiftResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cs/v2/cashmanagement/withdrawals/WithdrawIBanSwift")
	if err != nil {
		return nil, err
	}
	var respJson *WithdrawIBanSwiftResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type WithdrawIBanSwiftParams struct {
	AccountKey           string // required
	AdvancedOptions      string
	Amount               float64 // required
	ClearingCode         string
	Currency             string // required
	Iban                 string // required
	MessageToBeneficiary string
	ReceivingCountryCode string
	Swift                string
}

// https://www.developer.saxo/openapi/referencedocs/cs/v2/cashmanagement-cashwithdrawal/withdrawalaccountnumberswift/9a7fbe8d2d7a3817c47bbd35acbda78c
func WithdrawalAccountNumberSwift(params WithdrawalAccountNumberSwiftParams) (*WithdrawalAccountNumberSwiftResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cs/v2/cashmanagement/withdrawals/WithdrawalAccountNumberSwift")
	if err != nil {
		return nil, err
	}
	var respJson *WithdrawalAccountNumberSwiftResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type WithdrawalAccountNumberSwiftParams struct {
	AccountKey           string // required
	AccountNumber        string // required
	AdvancedOptions      string
	Amount               float64 // required
	ClearingCode         string
	Currency             string // required
	MessageToBeneficiary string
	ReceivingCountryCode string
	Swift                string
}

// https://www.developer.saxo/openapi/referencedocs/cs/v2/cashmanagement-cashwithdrawal/withdrawalaccountnumberbeneficiarybank/b68d380e739bc8b08a400af8e6c8c0f5
func WithdrawalAccountNumberBeneficiaryBank(params WithdrawalAccountNumberBeneficiaryBankParams) (*WithdrawalAccountNumberBeneficiaryBankResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cs/v2/cashmanagement/withdrawals/WithdrawalAccountNumberBeneficiaryBank")
	if err != nil {
		return nil, err
	}
	var respJson *WithdrawalAccountNumberBeneficiaryBankResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type WithdrawalAccountNumberBeneficiaryBankParams struct {
	AccountKey                 string // required
	AccountNumber              string // required
	AdvancedOptions            string
	Amount                     float64 // required
	BeneficiaryAddress         string  // required
	BeneficiaryCityAndPostcode string  // required
	BeneficiaryCountry         string  // required
	BeneficiaryName            string  // required
	ClearingCode               string
	Currency                   string // required
	MessageToBeneficiary       string
	ReceivingCountryCode       string
	Swift                      string
}

// https://www.developer.saxo/openapi/referencedocs/cs/v2/cashmanagement-cashwithdrawal/withdrawalibanbeneficiarybank/09eba5f22f167ed870b47e0328b20d73
func WithdrawalIBanBeneficiaryBank(params WithdrawalIBanBeneficiaryBankParams) (*WithdrawalIBanBeneficiaryBankResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cs/v2/cashmanagement/withdrawals/WithdrawalIbanBeneficiaryBank")
	if err != nil {
		return nil, err
	}
	var respJson *WithdrawalIBanBeneficiaryBankResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type WithdrawalIBanBeneficiaryBankParams struct {
	AccountKey                 string // required
	AdvancedOptions            string
	Amount                     float64 // required
	BeneficiaryAddress         string  // required
	BeneficiaryCityAndPostcode string  // required
	BeneficiaryCountry         string  // required
	BeneficiaryName            string  // required
	ClearingCode               string
	Currency                   string // required
	Iban                       string // required
	MessageToBeneficiary       string
	ReceivingCountryCode       string
	Swift                      string
}
