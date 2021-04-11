package autotrading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/getinvestments/06bbe2ae26a64d753b7d82de447046d0
func GetInvestments(params GetInvestmentsParams) (*GetInvestmentsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/investments/me/?TradeLeaderId={TradeLeaderId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetInvestmentsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetInvestmentsParams struct {
	TradeLeaderId string
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/getinvestment/4cad243b6c5a1d4d21f422bce85a4afb
func GetInvestment(params GetInvestmentParams) (*GetInvestmentResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/investments/?InvestmentId={InvestmentId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetInvestmentResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetInvestmentParams struct {
	InvestmentId string // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/getinvestmentbyleader/b1eeb498e18535a8849b79679024c8a3
func GetInvestmentByLeader(params GetInvestmentByLeaderParams) (*GetInvestmentByLeaderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/investments/?TradeLeaderId={TradeLeaderId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetInvestmentByLeaderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetInvestmentByLeaderParams struct {
	TradeLeaderId string // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/postinvestment/e0e757bd86e8433c2179f4691b529d0b
func PostInvestment(params PostInvestmentParams) (*PostInvestmentResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/at/v3/investments")
	if err != nil {
		return nil, err
	}
	var respJson *PostInvestmentResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type PostInvestmentParams struct {
	DisplayName            string
	FundingAccountKey      string  // required
	InitialFunding         float64 // required
	InvestmentShieldAmount float64 // required
	IsDisclaimerAccepted   bool    // required
	TradeLeaderId          string  // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/patchinvestmentasync/e9fb5b36287be5ae114e5ede06a38ca5
func PatchInvestmentAsync(investmentid string, params PatchInvestmentAsyncParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/at/v3/investments/{InvestmentId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type PatchInvestmentAsyncParams struct {
	DisplayName       string
	FundingAccountKey string

	InvestmentShieldAmount       float64
	IsActive                     bool
	IsAutoFundsTransfer          bool
	PendingFunding               float64
	PositionsCloseOnDeactivation bool
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/putinvestmentasync/52ab7fecb8ea06db0175f5498c2d8de9
func PutInvestmentAsync(investmentid string, params PutInvestmentAsyncParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/at/v3/investments/{InvestmentId}/withdraw")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type PutInvestmentAsyncParams struct {
	Amount float64 // required

	TargetAccountKey string // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/getmaxwithdrawalamountasync/3760b59eaa0a1408afdc165f5f59ff5d
func GetMaxWithdrawalAmountAsync(investmentid string, params GetMaxWithdrawalAmountAsyncParams) (*GetMaxWithdrawalAmountAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/investments/{InvestmentId}/maxWithdrawalAmount/?TargetAccountId={TargetAccountId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetMaxWithdrawalAmountAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetMaxWithdrawalAmountAsyncParams struct {
	TargetAccountId string // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/getinvestmentsummariesasync/aa74880cbb4f924210b0a8383a9c70ab
func GetInvestmentsummariesAsync() (*GetInvestmentsummariesAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/investments/summaries")
	if err != nil {
		return nil, err
	}
	var respJson *GetInvestmentsummariesAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/addsubscription/571d2982fe9d8bc638e70a595e009952
func AddSubscription(params AddSubscriptionParams) (*AddSubscriptionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/at/v3/investments/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionParams struct {
	Arguments   string
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/addsuggestionsubscriptionasync/a01be3aa98325261e7d4e8536aa37173
func AddSuggestionSubscriptionAsync(params AddSuggestionSubscriptionAsyncParams) (*AddSuggestionSubscriptionAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/at/v3/investments/subscriptions/suggestions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSuggestionSubscriptionAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSuggestionSubscriptionAsyncParams struct {
	Arguments   string
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/deletesubscription/886b909bc204cd913b761f20a9bfcbf2
func DeleteSubscription(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/at/v3/investments/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/investments/deletesuggestionsubscription/cd2ae815bc2766883a2dddc9ea5cca4c
func DeleteSuggestionSubscription(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/at/v3/investments/subscriptions/suggestions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
