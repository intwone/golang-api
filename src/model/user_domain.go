package model

import (
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/util"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func NewUserDomain(email string, password string, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hashedPassword, err := util.HashPassword(ud.Password)

	if err != nil {
		logger.Error("error during generate hash to password", err, util.CreateJourneyField("UserDomain"))
	}

	ud.Password = hashedPassword
}
