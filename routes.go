// routes.go

package main

func initializeRoutes() {

	router.Use(setUserStatus())
	//Handle the index route
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureNotLoggedIn(), logout)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}
