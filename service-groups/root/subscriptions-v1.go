package root

import (
	"github.com/cborum/go-saxo"
)

// https://www.developer.saxo/openapi/referencedocs/root/v1/subscriptions/deletesubscriptions/14d1d2d03f6aca45ab9c29e38973f449
func DeleteSubscriptions(contextid string, params DeleteSubscriptionsParams) ([]byte, error) {
	resp, err := saxo.GetClient().HttpClient.Delete("https://gateway.saxobank.com/sim/openapi/root/v1/subscriptions/{ContextId}/?Tag={Tag}")
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

type DeleteSubscriptionsParams struct {
	Tag string
}
