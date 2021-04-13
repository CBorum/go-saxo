// This file is autogenerated by cmd/generate/main.go
package trading

import (
    "fmt"

	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/trade/v1/allocationkeys/getallocationkeys/43a2f548b96be1dfa18b3a7cfaf6bb89
func GetAllocationKeys(params *GetAllocationKeysParams) (*GetAllocationKeysResponse, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/trade/v1/allocationkeys/?$top={$top}&$skip={$skip}&AccountKey={AccountKey}&ClientKey={ClientKey}&Statuses={Statuses}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &GetAllocationKeysResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type GetAllocationKeysParams struct { 
    skip int64
    top int64
    AccountKey string
    ClientKey string
    Statuses string 
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/allocationkeys/getallocationkey/6c3cf06c46f41128f215ed9f9a73a676
func GetAllocationKey(allocationkeyid string) (*GetAllocationKeyResponse, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/trade/v1/allocationkeys/{AllocationKeyId}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &GetAllocationKeyResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}


// https://www.developer.saxo/openapi/referencedocs/trade/v1/allocationkeys/distributeamountaccordingtoallocationkey/db4cff5278249a4e3d8b99f49690d5fb
func DistributeAmountAccordingToAllocationKey(allocationkeyid string, params *DistributeAmountAccordingToAllocationKeyParams) (*DistributeAmountAccordingToAllocationKeyResponse, error) {
    resp, err := saxo.GetClient().DoRequest("GET", "https://gateway.saxobank.com/sim/openapi/trade/v1/allocationkeys/distributions/{AllocationKeyId}/?Totalamount={Totalamount}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &DistributeAmountAccordingToAllocationKeyResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type DistributeAmountAccordingToAllocationKeyParams struct { 
    
    Totalamount string 
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/allocationkeys/addallocationkey/324a07c660d07971519f57d72bff1ea9
func AddAllocationKey(params *AddAllocationKeyParams) (*AddAllocationKeyResponse, error) {
    resp, err := saxo.GetClient().DoRequest("POST", "https://gateway.saxobank.com/sim/openapi/trade/v1/allocationkeys", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    respJson := &AddAllocationKeyResponse{}
    err = resp.ToJSON(respJson)
    if err != nil {
        return nil, err
    }
    return respJson, nil
}

type AddAllocationKeyParams struct { 
    AllocationKeyName string // required
    AllocationUnitType string // required
    MarginHandling string // required
    OneTime bool // required
    OwnerAccountKey string // required
    ParticipatingAccountsInfo string // required 
}

// https://www.developer.saxo/openapi/referencedocs/trade/v1/allocationkeys/deleteallocationkey/4186abc67c861a502ba1f26c9519dd1a
func DeleteAllocationKey(allocationkeyid string) ([]byte, error) {
    resp, err := saxo.GetClient().DoRequest("DELETE", "https://gateway.saxobank.com/sim/openapi/trade/v1/allocationkeys/{AllocationKeyId}", nil) 
    if err != nil {
        return nil, err
    } else if sc := resp.Response().StatusCode; sc >= 300 {
		return nil, fmt.Errorf("unexpected status code %d", sc)
	}
    return resp.Bytes(), nil 
}
