package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/configuration/validation"
	"github.com/intwone/golang-api/src/controller/model/request"
	"github.com/intwone/golang-api/src/controller/model/response"
)

func CreateUser(c *gin.Context) {
	var request request.UserRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		log.Printf("error trying to marshal object, error = %s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(request)

	response := response.UserResponse{
		Id:    "1234",
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}

	c.JSON(http.StatusOK, response)
}
