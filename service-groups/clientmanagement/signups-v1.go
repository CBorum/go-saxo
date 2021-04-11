package clientmanagement

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/attachfile/d4080612eac8c8aa989c8b2cd0cefe11
func AttachFilev1(signupid string, params AttachFilev1Params) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/attachments/{SignUpId}/?RenewalDate={RenewalDate}&DocumentType={DocumentType}&Title={Title}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type AttachFilev1Params struct {
	DocumentType string
	RenewalDate  string

	Title string
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/getsignupoptions/11227f361d27279ec9b7ce63abe47ecb
func GetSignupOptionsv1() (*GetSignupOptionsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/options")
	if err != nil {
		return nil, err
	}
	var respJson *GetSignupOptionsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/signup/e6279475c6da5c61867904c574531d2a
func SignUpv1(params SignUpv1Params) (*SignUpResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/?OwnerKey={OwnerKey}")
	if err != nil {
		return nil, err
	}
	var respJson *SignUpResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type SignUpv1Params struct {
	AccountInformation    string
	BankInformation       string
	FinlandData           string
	ItalyData             string
	OnboardingInformation string
	OwnerKey              string
	PensionData           string
	PersonalInformation   string // required
	ProfileInformation    string
	RegulatoryInformation string // required
	SingaporeData         string
	SwitzerlandData       string
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/getsignupstatus/41a070d0384723da2145a39ccd0bf9e0
func GetSignUpStatusv1(clientkey string) (*GetSignUpStatusResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/status/{ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetSignUpStatusResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/completeapplication/177fa8affca7cd529d28c888c697a5f4
func CompleteApplicationv1(signupid string, params CompleteApplicationv1Params) (*CompleteApplicationResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/completeapplication/{SignUpId}/?AwaitAccountCreation={AwaitAccountCreation}")
	if err != nil {
		return nil, err
	}
	var respJson *CompleteApplicationResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CompleteApplicationv1Params struct {
	AwaitAccountCreation bool
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/initiateverification/d19dc3a3d9ad7f366ae2f3f7583c6b99
func InitiateVerificationv1(clientkey string, params InitiateVerificationv1Params) (*InitiateVerificationResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/verification/initiate/{ClientKey}")
	if err != nil {
		return nil, err
	}
	var respJson *InitiateVerificationResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type InitiateVerificationv1Params struct {
	RedirectUrl string // required
}

// https://www.developer.saxo/openapi/referencedocs/cm/v1/signups/generatetypedonboardingpdf/6089ca2f4ae3a1af2d02c693b75ce1c9
func GenerateTypedOnboardingPDFv1(clientkey string, params GenerateTypedOnboardingPDFv1Params) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/cm/v1/signups/onboardingpdf/{ClientKey}/?DocumentType={DocumentType}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type GenerateTypedOnboardingPDFv1Params struct {
	DocumentType string // required
}
