package repository

import (
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepositoryInterface {
	return &userRepository{
		database: database,
	}
}
