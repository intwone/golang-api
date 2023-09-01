package model

import "github.com/intwone/golang-api/src/configuration/rest_err"

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetId(string)
	EncryptPassword()
	ComparePassword(password string, hashedPassword string) bool
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(email string, password string, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserSigInDomain(email string, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
