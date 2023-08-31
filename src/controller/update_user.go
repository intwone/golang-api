package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/configuration/validation"
	"github.com/intwone/golang-api/src/controller/model/request"
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	var request request.UserUpdateRequest

	userId := c.Param("userId")

	err := c.ShouldBindJSON(&request)

	if err != nil {
		message := fmt.Sprintf("error trying to validate user, error = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("UpdateUserController"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	_, objectIdErr := primitive.ObjectIDFromHex(userId)

	if objectIdErr != nil {
		restErr := rest_err.NewBadRequestError("Invalid user_id, must be a hex value")
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(request.Name, request.Age)

	serviceErr := uc.service.UpdateUser(userId, domain)

	if serviceErr != nil {
		message := fmt.Sprintf("error to UpdateUser in service, error = %s", serviceErr.Error())
		logger.Error(message, serviceErr, util.CreateJourneyField("UpdateUserController"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	c.Status(http.StatusOK)
}
