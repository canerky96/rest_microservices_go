package app

import "bookstore_users-api/controllers"

func mapUrls() {
	router.POST("/user", controllers.Create)
	router.GET("/user/:user_id", controllers.Get)
	router.GET("/user/search", controllers.Search)
}
