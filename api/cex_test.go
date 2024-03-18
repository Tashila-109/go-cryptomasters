package api_test

import (
	"cryptomasters/go/crypto/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Error was not found")
	}
}
