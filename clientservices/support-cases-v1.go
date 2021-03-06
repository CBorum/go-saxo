// This file is autogenerated by cmd/generate/main.go
package clientservices

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/getcases/9917b4984825a55600d2dc51547fd49e
func GetCases(params *GetCasesParams) (*GetCasesResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases/?$top={$top}&Status={Status}&FromDateTime={FromDateTime}&ToDateTime={ToDateTime}"
	url = saxo.PrepareUrlParams(url, params)
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetCasesResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type GetCasesParams struct {
	Top          int64  `url:",omitempty"`
	FromDateTime string `url:",omitempty"`
	Status       string `url:",omitempty"`
	ToDateTime   string `url:",omitempty"`
}

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/getcase/9cdeebac8805cd2800e1fd7373882475
func GetCase(caseid string) (*GetCaseResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases/{CaseId}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{CaseId}", caseid))
	resp, err := saxo.GetClient().DoRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	respJson := &GetCaseResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/closecase/166b34c2550c6861c293992c841cdef5
func CloseCase(caseid string, params *CloseCaseParams) ([]byte, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases/{CaseId}/caseclose"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{CaseId}", caseid))
	resp, err := saxo.GetClient().DoRequest("PUT", url, params)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type CloseCaseParams struct {
	Status string // required
}

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/createinternalcomment/74219517f6b3255e6dc59e91ecc4a456
func CreateInternalComment(caseid string, params *CreateInternalCommentParams) ([]byte, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases/{CaseId}/internalcomment"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{CaseId}", caseid))
	resp, err := saxo.GetClient().DoRequest("POST", url, params)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type CreateInternalCommentParams struct {
	Comment string // required
}

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/createnote/286e1f8e48dce194d7f3207edaa7b0ac
func CreateNote(caseid string, params *CreateNoteParams) ([]byte, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases/{CaseId}/note"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{CaseId}", caseid))
	resp, err := saxo.GetClient().DoRequest("POST", url, params)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type CreateNoteParams struct {
	Attachment string
	Note       string
	Title      string
}

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/updatenote/81ae6f1a1f9abdc0e39ca03adabc6dbd
func UpdateNote(caseid string, params *UpdateNoteParams) ([]byte, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases/{CaseId}"
	url = saxo.PrepareUrlRoute(url, saxo.RP("{CaseId}", caseid))
	resp, err := saxo.GetClient().DoRequest("PATCH", url, params)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type UpdateNoteParams struct {
	CaseStatus       string
	CaseType         string
	Description      string
	FollowUpDueDate  string
	HandledByPartner bool
	IsEscalated      bool
	ShowInPortal     bool
	Title            string
}

// https://www.developer.saxo/openapi/referencedocs/cs/v1/support-cases/createcase/136ad3589446a4b858a0d3f0c57c072e
func CreateCase(params *CreateCaseParams) (*CreateCaseResponse, error) {
	url := "https://gateway.saxobank.com/sim/openapi/cs/v1/partner/support/cases"
	resp, err := saxo.GetClient().DoRequest("POST", url, params)
	if err != nil {
		return nil, err
	}
	respJson := &CreateCaseResponse{}
	err = resp.ToJSON(respJson)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

type CreateCaseParams struct {
	CaseTitle    string // required
	ClientKey    string // required
	Description  string
	NotifyClient bool
	ShowInPortal bool
}
