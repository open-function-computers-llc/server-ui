package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	s := createDummyServer()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("We weren't expecting an error. Error: " + err.Error())
	}

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Incorrect status code.\nExpected: 200\n  Actual: %d", w.Code)
	}

	expectedInBody := "Welcome to the OfCO server maintainer"
	if !strings.Contains(w.Body.String(), expectedInBody) {
		t.Errorf("We expected to find the string `%s` in the response body. Response body is `%s`.", expectedInBody, w.Body)
	}
}
