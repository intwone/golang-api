package model

import (
	"fmt"

	"github.com/intwone/golang-api/src/configuration/rest_err"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
