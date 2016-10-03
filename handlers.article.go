//handlers.article.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Render a template via Context
	c.HTML(
		// Set http status to 200(OK)
		http.StatusOK,
		// Use index.html template
		"index.html",
		// Pass the page vars
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
