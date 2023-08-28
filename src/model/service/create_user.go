package service

import (
	"fmt"

	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
