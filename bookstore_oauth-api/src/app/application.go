package app

import (
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository"
	"bookstore_oauth-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := repository.NewRepository()
	atService := services.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	_ = router.Run(":8080")
}
