//handlers.user_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "HopFrog")
	params.Add("password", "1234")

	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/u/register", showRegistrationPage)

	req, _ := http.NewRequest("GET", "/u/register", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title") > 0

		return statusOK && pageOK
	})
}
