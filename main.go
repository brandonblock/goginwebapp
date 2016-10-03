package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set default router
	router = gin.Default

	// Process templates at initialization
	router.LoadHTMLGlob("templates/*")

	// Initialize routes
	initializeRoutes()

	// Start the application
	router.Run()
}
