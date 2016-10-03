package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// Set router with defaults
	router := gin.Default()
	// Process templates when app starts
	router.LoadHTMLGlob("templates/*")

	// Define route for index
	router.GET("/", func(c *gin.Context) {

		// Call the HTML Context method to render a template
		c.HTML(
			// Set the HTTP status to 200(OK)
			http.StatusOK,
			// Point to the index.html template
			"index.html",
			// Pass data to the template vars
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Start the server
	router.Run()
}
