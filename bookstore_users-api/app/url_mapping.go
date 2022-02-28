package app

import "bookstore_users-api/controllers"

func mapUrls() {
	router.POST("/user", controllers.Create)
	router.GET("/user/:user_id", controllers.Get)
	router.PUT("/user/:user_id", controllers.UpdateUser)
	router.PATCH("/user/:user_id", controllers.PartialUpdateUser)
	router.GET("/user/search", controllers.Search)
}
