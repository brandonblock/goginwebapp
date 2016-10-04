// routes.go

package main

func initializeRoutes() {
	//Handle the index route
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/")
	{
		userRoutes.GET("/register", showRegistrationPage)
		userRoutes.POST("/register", register)
	}

	//Handle requests for single articles
	router.GET("/article/view/:article_id", getArticle)
}
