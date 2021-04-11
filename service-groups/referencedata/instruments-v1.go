package referencedata

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/ref/v1/instruments/getsummaries/3c16fe9480ce291643f09a437cfc001d
func GetSummaries(params GetSummariesParams) (*GetSummariesResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/instruments/?$top={$top}&$skip={$skip}&ExchangeId={ExchangeId}&Keywords={Keywords}&SectorId={SectorId}&IncludeNonTradable={IncludeNonTradable}&Class={Class}&CanParticipateInMultiLegOrder={CanParticipateInMultiLegOrder}&Uics={Uics}&AssetTypes={AssetTypes}&Tags={Tags}&AccountKey={AccountKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetSummariesResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetSummariesParams struct {
	skip                          int64
	top                           int64
	AccountKey                    string
	AssetTypes                    string
	CanParticipateInMultiLegOrder bool
	Class                         string
	ExchangeId                    string
	IncludeNonTradable            bool
	Keywords                      string
	SectorId                      string
	Tags                          string
	Uics                          string
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/instruments/getdetailsformanyinstruments/ad9c80ea6ddc7c7974d653e45a495f87
func GetDetailsForManyInstruments(params GetDetailsForManyInstrumentsParams) (*GetDetailsForManyInstrumentsResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/instruments/details/?$top={$top}&$skip={$skip}&Uics={Uics}&AssetTypes={AssetTypes}&Tags={Tags}&AccountKey={AccountKey}&FieldGroups={FieldGroups}")
	if err != nil {
		return nil, err
	}
	var respJson *GetDetailsForManyInstrumentsResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetDetailsForManyInstrumentsParams struct {
	skip        int64
	top         int64
	AccountKey  string
	AssetTypes  string
	FieldGroups string
	Tags        string
	Uics        string
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/instruments/getdetailsforoneinstrument/4e547901bf01a81683b01ee9ed9aca0b
func GetDetailsForOneInstrument(assettype string, uic string, params GetDetailsForOneInstrumentParams) (*GetDetailsForOneInstrumentResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/instruments/details/{Uic}/{AssetType}/?FieldGroups={FieldGroups}&AccountKey={AccountKey}&ClientKey={ClientKey}&MarketDataProvider={MarketDataProvider}&IgnorePermission={IgnorePermission}")
	if err != nil {
		return nil, err
	}
	var respJson *GetDetailsForOneInstrumentResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetDetailsForOneInstrumentParams struct {
	AccountKey string

	ClientKey   string
	FieldGroups string
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/instruments/getcontractoptionspace/9168dcd1c267be2c577576179daaa2fb
func GetContractOptionSpace(optionrootid string, params GetContractOptionSpaceParams) (*GetContractOptionSpaceResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/instruments/contractoptionspaces/{OptionRootId}/?ClientKey={ClientKey}&OptionSpaceSegment={OptionSpaceSegment}&ExpiryDates={ExpiryDates}&UnderlyingUic={UnderlyingUic}&CanParticipateInMultiLegOrder={CanParticipateInMultiLegOrder}&TradingStatus={TradingStatus}")
	if err != nil {
		return nil, err
	}
	var respJson *GetContractOptionSpaceResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetContractOptionSpaceParams struct {
	CanParticipateInMultiLegOrder bool
	ClientKey                     string
	ExpiryDates                   string

	OptionSpaceSegment string
	TradingStatus      string
	UnderlyingUic      int64
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/instruments/getfuturesspace/ab6c53dae08529ade5f6ab7b1aa7f27a
func GetFuturesSpace(continuousfuturesuic string) (*GetFuturesSpaceResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/instruments/futuresspaces/{ContinuousFuturesUic}")
	if err != nil {
		return nil, err
	}
	var respJson *GetFuturesSpaceResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/ref/v1/instruments/gettradingschedule/6ae566f8ab4ce0b3cb4b1578d5d5449c
func GetTradingSchedule(assettype string, uic string) (*GetTradingScheduleResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/ref/v1/instruments/tradingschedule/{Uic}/{AssetType}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradingScheduleResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}
