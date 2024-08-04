package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/:userId")
	r.GET("/users/email/:email")
	r.POST("/users")
	r.PUT("/users/:userId")
	r.DELETE("users/:userId")
}
