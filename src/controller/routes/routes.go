package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ruhtar/webapi-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/:userId", controller.FindUserById)
	r.GET("/users/email/:email", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("users/:userId", controller.DeleteUser)
}
