package autotrading

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/getalltradeleaders/a591321e9eabc01123207a1aa0ab0a76
func GetAllTradeLeaders(params GetAllTradeLeadersParams) (*GetAllTradeLeadersResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/?StandardPeriod={StandardPeriod}")
	if err != nil {
		return nil, err
	}
	var respJson *GetAllTradeLeadersResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAllTradeLeadersParams struct {
	StandardPeriod string
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/gettradeleader/1a253bc4aaca7985e3ee5341673d75aa
func GetTradeLeader(tradeleaderid string, params GetTradeLeaderParams) (*GetTradeLeaderResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/{TradeLeaderId}/?StandardPeriod={StandardPeriod}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradeLeaderResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetTradeLeaderParams struct {
	StandardPeriod string
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/gettradeleaderlogo/c9ea3fcfebf11c21337a0ddc204116a9
func GetTradeLeaderLogo(tradeleaderid string) (*GetTradeLeaderLogoResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/{TradeLeaderId}/logo")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradeLeaderLogoResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/gettradeleaderlogobylogokey/4db243e57a760261d93e2242eace088f
func GetTradeLeaderLogoByLogoKey(logokey string) (*GetTradeLeaderLogoByLogoKeyResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/logo/{LogoKey}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradeLeaderLogoByLogoKeyResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/getalltradeleaderspublic/8a13092fa9bdc32c7e257bc643e7095e
func GetAllTradeLeadersPublic(params GetAllTradeLeadersPublicParams) (*GetAllTradeLeadersPublicResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/public")
	if err != nil {
		return nil, err
	}
	var respJson *GetAllTradeLeadersPublicResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetAllTradeLeadersPublicParams struct {
	OwnerKey       string
	StandardPeriod string
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/gettradeleaderpublic/61e636010344e6034bbf62f3f31c2f54
func GetTradeLeaderPublic(params GetTradeLeaderPublicParams) (*GetTradeLeaderPublicResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/public/{TradeLeaderId}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradeLeaderPublicResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetTradeLeaderPublicParams struct {
	OwnerKey       string
	StandardPeriod string
	TradeLeaderId  int64 // required
}

// https://www.developer.saxo/openapi/referencedocs/at/v3/tradeleaders/gettradeleaderscustomselectionpublic/14cee38313eacabb2b36a4dd76ea0dd0
func GetTradeLeadersCustomSelectionPublic(selectionid string, params GetTradeLeadersCustomSelectionPublicParams) (*GetTradeLeadersCustomSelectionPublicResponse, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/at/v3/tradeLeaders/public/customselection/{SelectionId}/?StandardPeriod={StandardPeriod}")
	if err != nil {
		return nil, err
	}
	var respJson *GetTradeLeadersCustomSelectionPublicResponse
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetTradeLeadersCustomSelectionPublicParams struct {
	StandardPeriod string
}
