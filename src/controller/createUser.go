package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ruhtar/webapi-go/src/configuration/rest_err"
	"github.com/ruhtar/webapi-go/src/controller/dtos/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_err := rest_err.NewBadRequestError(fmt.Sprintf("Incorrect fields, erro=%s", err.Error()))
		c.JSON(rest_err.Code, rest_err)
		return
	}

	c.JSON(200, userRequest)
}
