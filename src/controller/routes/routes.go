package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/id/:userId", controller.FindUserById)
	r.GET("/user/email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
