package app

import (
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := repository.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	_ = router.Run(":8080")
}
