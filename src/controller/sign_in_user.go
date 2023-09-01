package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/configuration/validation"
	"github.com/intwone/golang-api/src/controller/model/request"
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/util"
	"github.com/intwone/golang-api/src/view"
)

func (uc *userControllerInterface) SignInUser(c *gin.Context) {
	var request request.UserLoginRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		message := fmt.Sprintf("error trying to marshal object, error = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("SignInUserController"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserSigInDomain(request.Email, request.Password)

	domainResult, serviceErr := uc.service.SignInUser(domain)

	if serviceErr != nil {
		message := fmt.Sprintf("error to SignIn in service, error = %s", serviceErr.Error())
		logger.Error(message, serviceErr, util.CreateJourneyField("SignInUserController"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	response := view.ConvertDomainToResponse(domainResult)

	c.JSON(http.StatusOK, response)
}
