package main

import (
	"fmt"
	"testing"
)

func TestGgetResponseBodyStruct(t *testing.T) {
	bb, err := getResponseBodyStruct("https://www.developer.saxo/openapi/referencedocs/ref/v1/algostrategies/getstrategybyname/f3195d507d54dd21c5504df61bf17863", "GetStrategyByName", "algostrategies")
	if err != nil {
		t.FailNow()
	}
	fmt.Println(string(bb))

	if len(bb) == 0 {
		t.Fail()
	}
}

func TestGgetResponseBodyStruct2(t *testing.T) {
	bb, err := getResponseBodyStruct("https://www.developer.saxo/openapi/referencedocs/vas/v1/pricealerts/putalertdefinition/5e45019d26e97ad5e7e5ec17f107982b", "GetStrategyByName", "algostrategies")
	if err != nil {
		t.FailNow()
	}
	fmt.Println(string(bb))

	if len(bb) > 0 {
		t.Fail()
	}
}
