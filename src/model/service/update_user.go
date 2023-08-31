package service

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
)

func (ud *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	err := ud.repository.UpdateUser(id, userDomain)

	if err != nil {
		return err
	}

	return nil
}
