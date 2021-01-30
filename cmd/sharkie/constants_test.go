package main

import (
	"testing"
)

func TestSetHeaders(t *testing.T) {
	// check to make sure space is being trimmed appropriately
	x := []string{"   test   :   header value  "}
	expectedValue := "header value"
	setheaders(x)

	if TDATA.Headers["test"] != expectedValue {
		t.Errorf("Header value: \"%v\" - Not trimming spaces correctly. Expected \"%v\"", TDATA.Headers["test"], expectedValue)
	}
}
