package service

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	err := ud.repository.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
