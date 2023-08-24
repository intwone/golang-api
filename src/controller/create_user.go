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
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
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

	domainErr := domain.CreateUser()

	if domainErr != nil {
		message := fmt.Sprintf("error to CreateUser in domain, error = %s", domainErr.Error())
		logger.Error(message, err, util.CreateJourneyField("CreateUserController"))
		c.JSON(domainErr.Code, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
