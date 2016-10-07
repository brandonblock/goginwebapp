// models.user_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUsernameAvailability(t *testing.T) {
	saveLists()

	if !isUsernameAvailable("Frodo") {
		t.Fail()
	}

	if isUsernameAvailable("HopFrog") {
		t.Fail()
	}

	registerNewUser("newuser", "newpass")

	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("HopFrog", "1234")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUserValidity(t *testing.T) {
	if !isUserValid("HopFrog", "1234") {
		t.Fail()
	}

	if isUserValid("GoldBug", "1234") {
		t.Fail()
	}

	if isUserValid("HopFrog", "") {
		t.Fail()
	}

	if isUserValid("hopFrog", "1234") {
		t.Fail()
	}
}

func TestShowLoginPageUnathenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/u/login", showLoginPage)

	req, _ := http.NewRequest("GET", "/u/login", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

func TestLoginUnauthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.POST("/u/login", performLogin)

	loginPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful Login</title>") < 0 {
		t.Fail()
	}
	restoreLists()
}

func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.POST("/u/login", performLogin)

	loginPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
	restoreLists()
}
