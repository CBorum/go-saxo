package trading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v1/positions/createposition/d3387419f4ffb6c9877997702cb75ae2
func CreatePosition(params CreatePositionParams) (*CreatePositionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/trade/v1/positions")
	if err != nil {
		return nil, err
	}
	var respJson *CreatePositionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CreatePositionParams struct {
	AppHint          int64
	BuySell          string // required
	ContextId        string // required
	ExerciseMethod   string
	PriceReferenceId string // required
	QuoteId          string // required
	TradeIdeaId      int64
	UserPrice        float64 // required
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/positions/updateposition/dc09651b2b70d8fbd41ffab64e335cb4
func UpdatePosition(positionid string, params UpdatePositionParams) (*UpdatePositionResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/trade/v1/positions/{PositionId}")
	if err != nil {
		return nil, err
	}
	var respJson *UpdatePositionResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type UpdatePositionParams struct {
	AccountKey       string // required
	DisableForceOpen bool
	ExerciseMethod   string
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/positions/exercisepositionasync/91ec1719ac32f69dd30414ced6876414
func ExercisePositionAsync(positionid string, params ExercisePositionAsyncParams) (*ExercisePositionAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/trade/v1/positions/{PositionId}/exercise")
	if err != nil {
		return nil, err
	}
	var respJson *ExercisePositionAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type ExercisePositionAsyncParams struct {
	AccountKey string // required
	Amount     int64  // required
	AppHint    int64
	AssetType  string // required

	Uic int64 // required
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/positions/exerciseamountasync/98678d03e1c50ee18e3b50d80d43df88
func ExerciseAmountAsync(params ExerciseAmountAsyncParams) (*ExerciseAmountAsyncResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/trade/v1/positions/exercise")
	if err != nil {
		return nil, err
	}
	var respJson *ExerciseAmountAsyncResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type ExerciseAmountAsyncParams struct {
	AccountKey string // required
	Amount     int64  // required
	AppHint    int64
	AssetType  string // required
	Uic        int64  // required
}
