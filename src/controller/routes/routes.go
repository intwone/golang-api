package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/user/id/:userId", userController.FindUserById)
	r.GET("/user/email/:userEmail", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}
