package partnerintegration

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/partnerintegration/v1/interactiveidverification/idnowinteractiveverificationcomplete/7e8381ef50e281863d8fa9b1a76b8fb3
func IdNowInteractiveVerificationComplete(params IdNowInteractiveVerificationCompleteParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/partnerintegration/v1/InteractiveIdVerification/idnow/complete")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type IdNowInteractiveVerificationCompleteParams struct {
	VerificationStatus string
}

// https://www.developer.saxo/openapi/referencedocs/partnerintegration/v1/interactiveidverification/verifyinteractiveverification/409150c2211b11cdcc88e9125cc604d6
func VerifyInteractiveVerification(correlationid string, params VerifyInteractiveVerificationParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/partnerintegration/v1/InteractiveIdVerification/verify/{Correlationid}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type VerifyInteractiveVerificationParams struct {
	RequestBody string
}
