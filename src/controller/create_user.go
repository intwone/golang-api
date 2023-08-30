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

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var request request.UserRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		message := fmt.Sprintf("error trying to marshal object, error = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("CreateUserController"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(request.Email, request.Password, request.Name, request.Age)

	domainResult, serviceErr := uc.service.CreateUser(domain)

	if serviceErr != nil {
		message := fmt.Sprintf("error to CreateUser in service, error = %s", serviceErr.Error())
		logger.Error(message, serviceErr, util.CreateJourneyField("CreateUserController"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	response := view.ConvertDomainToResponse(domainResult)

	c.JSON(http.StatusCreated, response)
}
