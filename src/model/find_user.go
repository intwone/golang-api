package model

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
)

func (*UserDomain) FindUser(id string) (*UserDomain, *rest_err.RestErr) {
	return &UserDomain{
		Email:    "",
		Password: "",
		Name:     "",
		Age:      1,
	}, nil
}
