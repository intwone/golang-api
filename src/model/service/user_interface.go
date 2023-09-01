package service

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/model/repository"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
	SignInUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}

type userDomainService struct {
	repository repository.UserRepositoryInterface
}

func NewUserDomainService(repository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{
		repository: repository,
	}
}
