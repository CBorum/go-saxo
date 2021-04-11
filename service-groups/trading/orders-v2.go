package trading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/placeorder/33c53d35579766ba5243c8b5dd246e1a
func PlaceOrder(params PlaceOrderParams) (*PlaceOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v2/orders")
	if err != nil {
		return nil, err
	}
	var respJson *PlaceOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type PlaceOrderParams struct {
	AccountKey                   string
	AlgoOrderData                string
	AllocationKeyId              string
	Amount                       float64
	AmountType                   string
	AppHint                      int64
	AssetType                    string
	BuySell                      string
	CancelOrders                 bool
	ClearForceOpen               bool
	ExternalReference            string
	ForwardDate                  string
	IsForceOpen                  bool
	ManualOrder                  bool
	OrderDuration                string
	OrderId                      string
	OrderPrice                   float64
	Orders                       string
	OrderType                    string
	PositionId                   string
	QuoteCurrency                bool
	StopLimitPrice               float64
	SwitchInstrumentUic          int64
	ToOpenClose                  string
	TraderId                     string
	TrailingStopDistanceToMarket float64
	TrailingStopStep             float64
	TraspasoIn                   string
	Uic                          int64
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/changeorder/08fb9ab64437857d43a7b523e5149354
func ChangeOrder(params ChangeOrderParams) (*ChangeOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/trade/v2/orders")
	if err != nil {
		return nil, err
	}
	var respJson *ChangeOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type ChangeOrderParams struct {
	AccountKey                   string
	AlgoOrderData                string
	Amount                       float64
	AssetType                    string
	IsForceOpen                  bool
	OrderDuration                string
	OrderId                      string
	OrderPrice                   float64
	Orders                       string
	OrderType                    string
	StopLimitPrice               float64
	TraderId                     string
	TrailingStopDistanceToMarket float64
	TrailingStopStep             float64
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/cancelorder/a1fd2fffa62f21901f23318f65fe8147
func CancelOrder(orderids string, params CancelOrderParams) (*CancelOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/{OrderIds}/?AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *CancelOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CancelOrderParams struct {
	AccountKey string // required

}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/cancelallorder/261fb62628b4c2cc9278b170429c16fb
func CancelAllOrder(params CancelAllOrderParams) (*CancelAllOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/?AccountKey={AccountKey}&AssetType={AssetType}&Uic={Uic}")
	if err != nil {
		return nil, err
	}
	var respJson *CancelAllOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CancelAllOrderParams struct {
	AccountKey string // required
	AssetType  string // required
	Uic        int64  // required
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/precheckorder/a6278b748802edbefb73e5ccf9b30673
func PreCheckOrder(params PreCheckOrderParams) (*PreCheckOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/precheck")
	if err != nil {
		return nil, err
	}
	var respJson *PreCheckOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type PreCheckOrderParams struct {
	AccountKey          string  // required
	Amount              float64 // required
	AmountType          string
	AssetType           string // required
	BuySell             string // required
	ExternalReference   string
	FieldGroups         string
	ForwardDate         string
	IsForceOpen         bool
	OrderDuration       string
	OrderId             string
	OrderPrice          float64
	OrderType           string // required
	PositionId          string
	StopLimitPrice      float64
	StrategyId          string
	SwitchInstrumentUic int64
	ToOpenClose         string
	TraspasoIn          string
	Uic                 int64 // required
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/precheckmultilegorder/2495fa5bb96132fc94c912f4ebfa0d9d
func PreCheckMultilegOrder(params PreCheckMultilegOrderParams) (*PreCheckMultilegOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/multileg/precheck")
	if err != nil {
		return nil, err
	}
	var respJson *PreCheckMultilegOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type PreCheckMultilegOrderParams struct {
	AccountKey        string
	ExternalReference string
	FieldGroups       string
	Legs              string // required
	ManualOrder       bool
	OrderDuration     string
	OrderPrice        float64
	OrderType         string // required
	TraderId          string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/placemultilegstrategyorder/9a0a66f71c51c3320dd7baafdd179c25
func PlaceMultiLegStrategyOrder(params PlaceMultiLegStrategyOrderParams) (*PlaceMultiLegStrategyOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/multileg")
	if err != nil {
		return nil, err
	}
	var respJson *PlaceMultiLegStrategyOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type PlaceMultiLegStrategyOrderParams struct {
	AccountKey        string
	ExternalReference string
	Legs              string // required
	ManualOrder       bool
	OrderDuration     string
	OrderPrice        float64
	OrderType         string // required
	TraderId          string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/changemultilegstrategyorder/bcf9478928c21dd91e6ac1c44488a8af
func ChangeMultiLegStrategyOrder(params ChangeMultiLegStrategyOrderParams) (*ChangeMultiLegStrategyOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/multileg")
	if err != nil {
		return nil, err
	}
	var respJson *ChangeMultiLegStrategyOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type ChangeMultiLegStrategyOrderParams struct {
	AccountKey      string // required
	MultiLegAmount  float64
	MultiLegOrderId string // required
	OrderPrice      float64
	TraderId        string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/cancelmultilegstrategyorder/b4c8059c1ea08e1e65c06b75f76f516c
func CancelMultiLegStrategyOrder(multilegorderid string, params CancelMultiLegStrategyOrderParams) (*CancelMultiLegStrategyOrderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/multileg/{MultiLegOrderId}/?AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *CancelMultiLegStrategyOrderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CancelMultiLegStrategyOrderParams struct {
	AccountKey string // required

}

// https://www.developer.saxo/openapi/referencedocs/trade/v2/orders/getmultilegorderstrategydefaults/ec2e10daa55d46a3e3f1a57339d936d6
func GetMultiLegOrderStrategyDefaults(params GetMultiLegOrderStrategyDefaultsParams) (*GetMultiLegOrderStrategyDefaultsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/trade/v2/orders/multileg/defaults/?AccountKey={AccountKey}&OptionRootId={OptionRootId}&OptionsStrategyType={OptionsStrategyType}")
	if err != nil {
		return nil, err
	}
	var respJson *GetMultiLegOrderStrategyDefaultsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetMultiLegOrderStrategyDefaultsParams struct {
	AccountKey          string // required
	OptionRootId        int64  // required
	OptionsStrategyType string // required
}
