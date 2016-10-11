//middleware.auth.go

package main

import (
	"github.com/gin-gonic/gin"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
