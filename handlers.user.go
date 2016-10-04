//handlers.user.go

package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Fake insecure token generation
func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

// Simply call render function
func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	// Obtain POSTed user/pass values
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Set the token in a cookie if user succesfully created
	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Succesful registration & Login"}, "login-succesful.html")
	} else {
		// Show errors
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}
