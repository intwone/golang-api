package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/util"
	"github.com/intwone/golang-api/src/view"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	_, err := uuid.Parse(userId)

	if err != nil {
		message := fmt.Sprintf("userId is not a valid id")
		logger.Error(message, err, util.CreateJourneyField("FindUserByIdController"))
		errorMessage := rest_err.NewBadRequestError(message)
		c.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, serviceErr := uc.service.FindUserById(userId)

	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	response := view.ConvertDomainToResponse(userDomain)

	c.JSON(http.StatusOK, response)
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")

	_, err := mail.ParseAddress(userEmail)

	if err != nil {
		message := fmt.Sprintf("email is not a valid email")
		logger.Error(message, err, util.CreateJourneyField("FindUserByEmailController"))
		errorMessage := rest_err.NewBadRequestError(message)
		c.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, serviceErr := uc.service.FindUserByEmail(userEmail)

	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	response := view.ConvertDomainToResponse(userDomain)

	c.JSON(http.StatusOK, response)
}
