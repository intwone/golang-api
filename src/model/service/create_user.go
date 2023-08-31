package service

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	user, _ := ud.FindUserByEmail(userDomain.GetEmail())

	if user != nil {
		return nil, rest_err.NewBadRequestError("email already taken")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.repository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
