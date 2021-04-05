package server

import "testing"

func TestServerHostTypeChecking(t *testing.T) {
	err := verifyHostType("")
	if err == nil {
		t.Error("A default string isn't have a valid host type")
	}

	err = verifyHostType("modx")
	if err != nil {
		t.Error("'modx' is a valid host type")
	}
}
