package main

import (
	"github.com/intwone/golang-api/src/controller"
	"github.com/intwone/golang-api/src/model/repository"
	"github.com/intwone/golang-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repository)
	return controller.NewUserControllerInterface(service)
}
