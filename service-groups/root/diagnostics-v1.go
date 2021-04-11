package root

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/get/587ecb1c8e6d7c9c00749c08df3af0d1
func GetDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Get("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/get")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/post/b604ce168031c07f259dd84a07dc8f41
func PostDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Post("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/post")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/put/9bc2a0831d9e9df85c126e4ff96348a7
func PutDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Put("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/put")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/delete/248af0cb95c5a8e36fecf18d532a9511
func DeleteDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/delete")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/patch/9170f1a38321249d883d86003513aa6a
func PatchDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Patch("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/patch")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/head/144a22a97bc687397510353f11cbc2af
func HeadDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Head("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/head")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

// https://www.developer.saxo/openapi/referencedocs/root/v1/diagnostics/options/499592fb36c5e14d5712f78955301aca
func OptionsDiagnostics() ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Options("https://gateway.saxobank.com/sim/openapi/root/v1/diagnostics/options")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
