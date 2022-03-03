package app

import (
	"bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.POST("/user", users.Create)
	router.GET("/user/:user_id", users.Get)
	router.PUT("/user/:user_id", users.Update)
	router.PATCH("/user/:user_id", users.PartialUpdate)
	router.DELETE("/user/:user_id", users.Delete)
	router.GET("/user/search", users.Search)
}
