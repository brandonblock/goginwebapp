//middleware.auth_test.go

package main

import (
	"net/http"
)

func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggindIn(false), ensureLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}
