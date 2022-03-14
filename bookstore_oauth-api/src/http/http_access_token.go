package http

import (
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/services"
	"bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service services.Service
}

func NewHandler(service services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var token access_token.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := handler.service.Create(token); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
