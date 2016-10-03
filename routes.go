// routes.go

package main

func initializeRoutes() {
	//Handle the index route
	router.GET("/", showIndexPage)

	//Handle requests for single articles
	router.GET("/article/view/:articles_id", getArticle)
}
