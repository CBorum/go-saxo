package trading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v1/infoprices/getinfopriceasync/eee3cc82474270ca79836aa7d8b1e923
func GetInfoPriceAsync(params GetInfoPriceAsyncParams) (*GetInfoPriceAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/trade/v1/infoprices/?Uic={Uic}&AccountKey={AccountKey}&AssetType={AssetType}&Amount={Amount}&ForwardDate={ForwardDate}&ExpiryDate={ExpiryDate}&StrikePrice={StrikePrice}&OrderAskPrice={OrderAskPrice}&OrderBidPrice={OrderBidPrice}&LowerBarrier={LowerBarrier}&UpperBarrier={UpperBarrier}&PutCall={PutCall}&FieldGroups={FieldGroups}&AmountType={AmountType}&ToOpenClose={ToOpenClose}&QuoteCurrency={QuoteCurrency}")
	if err != nil {
		return nil, err
	}
	var respJson *GetInfoPriceAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetInfoPriceAsyncParams struct {
	AccountKey    string
	Amount        float64
	AmountType    string
	AssetType     string // required
	ExpiryDate    string
	FieldGroups   string
	ForwardDate   string
	LowerBarrier  float64
	OrderAskPrice float64
	OrderBidPrice float64
	PutCall       string
	QuoteCurrency bool
	StrikePrice   float64
	ToOpenClose   string
	Uic           int64 // required
	UpperBarrier  float64
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/infoprices/getinfopricelistasync/2eaaceb6373a7eff36c5f04f345cabe0
func GetInfoPriceListAsync(params GetInfoPriceListAsyncParams) (*GetInfoPriceListAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/trade/v1/infoprices/list/?Uics={Uics}&AccountKey={AccountKey}&AssetType={AssetType}&Amount={Amount}&ForwardDate={ForwardDate}&ExpiryDate={ExpiryDate}&StrikePrice={StrikePrice}&OrderAskPrice={OrderAskPrice}&OrderBidPrice={OrderBidPrice}&LowerBarrier={LowerBarrier}&UpperBarrier={UpperBarrier}&PutCall={PutCall}&FieldGroups={FieldGroups}&AmountType={AmountType}&ForwardDateNearLeg={ForwardDateNearLeg}&ForwardDateFarLeg={ForwardDateFarLeg}&ToOpenClose={ToOpenClose}&QuoteCurrency={QuoteCurrency}")
	if err != nil {
		return nil, err
	}
	var respJson *GetInfoPriceListAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetInfoPriceListAsyncParams struct {
	AccountKey         string
	Amount             float64
	AmountType         string
	AssetType          string // required
	ExpiryDate         string
	FieldGroups        string
	ForwardDate        string
	ForwardDateFarLeg  string
	ForwardDateNearLeg string
	LowerBarrier       float64
	OrderAskPrice      float64
	OrderBidPrice      float64
	PutCall            string
	QuoteCurrency      bool
	StrikePrice        float64
	ToOpenClose        string
	Uics               string // required
	UpperBarrier       float64
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/infoprices/addsubscriptionasync/38ca8a186fa551f5c2e16e9d7a25a7e2
func AddSubscriptionAsyncInfoPrices(params AddSubscriptionAsyncInfoPricesParams) (*AddSubscriptionAsyncInfoPricesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/infoprices/subscriptions")
	if err != nil {
		return nil, err
	}
	var respJson *AddSubscriptionAsyncInfoPricesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type AddSubscriptionAsyncInfoPricesParams struct {
	Arguments   string // required
	ContextId   string // required
	Format      string
	ReferenceId string // required
	RefreshRate int64
	Tag         string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/infoprices/deletesubscriptions/ba6b3598c3f0315f354250d65a8afb86
func DeleteSubscriptionsInfoPrices(contextid string, params DeleteSubscriptionsInfoPricesParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/infoprices/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsInfoPricesParams struct {
	Tag string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/infoprices/deletesubscription/f0c8e65d8cf01479377d1d81786b9ab4
func DeleteSubscriptionInfoPrices(contextid string, referenceid string) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v1/infoprices/subscriptions/{ContextId}/{ReferenceId}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
