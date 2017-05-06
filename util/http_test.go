package util

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJson(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"key":"value"}`)
	}))
	defer ts.Close()

	var result = &testDomainObj{}

	err := GetJson(ts.URL, result)

	if result.Key != "value" {
		t.Error("Failed to get json into domain object with error: ", err)
	}

	if err != nil {
		t.Error("Failed to perform HTTP GET with error: ", err)
	}

}

type testDomainObj struct {
	Key string
}
