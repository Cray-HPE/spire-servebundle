package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// TestTestRequestBundle verifies that the http server returns the expected bundle using
// environmental variables
func TestRequestBundle(t *testing.T) {
	// Bundle contains the struct used to generate the spire info json
	type Bundle struct {
		Domain     string
		Server     string
		CertBundle string
	}

	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("DOMAIN", "TEST.DOMAIN")
	os.Setenv("SERVER", "TEST.SERVER")
	os.Setenv("BUNDLE", path+"/testbundle.crt")

	r := strings.NewReader("")
	req, err := http.NewRequest("GET", "/", r)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rr, req)

	// Validate that Content-Type is properly set
	actualCT := rr.Header().Get("Content-Type")
	expectedCT := "application/json; charset=UTF-8"

	if actualCT != expectedCT {
		t.Fail()
		t.Logf("Receved wrong Content-Type\nExpected: %+v\n Actual: %+v", expectedCT, actualCT)
	}

	// Validate JSON
	expectedJSON := Bundle{"TEST.DOMAIN", "TEST.SERVER", "----- TEST -----\nTHIS IS A TEST FILE\n----- TEST -----\n"}
	var actualJSON Bundle
	body, err := ioutil.ReadAll(rr.Body)
	json.Unmarshal(body, &actualJSON)
	if actualJSON != expectedJSON {
		t.Fail()
		t.Logf("Received wrong JSON response\nExpected:\n%+v\n Actual:\n%+v\n", expectedJSON, actualJSON)
	}

}
