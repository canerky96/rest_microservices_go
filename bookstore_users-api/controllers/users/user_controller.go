package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}
func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	user.Id = userId
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func PartialUpdate(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	user.Id = userId
	result, err := services.PartialUpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {}