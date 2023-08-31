package service

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
)

func (ud *userDomainService) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ud.repository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ud.repository.FindUserByEmail(email)
}
