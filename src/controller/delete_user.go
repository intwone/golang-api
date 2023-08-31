package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	_, objectIdErr := primitive.ObjectIDFromHex(userId)

	if objectIdErr != nil {
		restErr := rest_err.NewBadRequestError("Invalid user_id, must be a hex value")
		c.JSON(restErr.Code, restErr)
		return
	}

	serviceErr := uc.service.DeleteUser(userId)

	if serviceErr != nil {
		message := fmt.Sprintf("error to DeleteUser in service, error = %s", serviceErr.Error())
		logger.Error(message, serviceErr, util.CreateJourneyField("DeleteUserController"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	c.Status(http.StatusOK)
}
