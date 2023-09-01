package service

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
)

func (ud *userDomainService) SignInUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	user, err := ud.FindUserByEmail(userDomain.GetEmail())

	if err != nil {
		return nil, "", rest_err.NewBadRequestError("email or password invalid")
	}

	compare := userDomain.ComparePassword(userDomain.GetPassword(), user.GetPassword())

	if !compare {
		return nil, "", rest_err.NewBadRequestError("email or password invalid")
	}

	token, tokenGenerateErr := user.GenerateToken()

	if tokenGenerateErr != nil {
		return nil, "", tokenGenerateErr
	}

	return user, token, nil
}
