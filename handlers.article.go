//handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Render the template
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check validity of article ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Render response
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")
		} else {
			// If article not found, abort with error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
