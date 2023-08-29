package repository

import (
	"context"

	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.database.Collection("users")

	value := converter.ConverterDomainToEntity(userDomain)

	result, insertErr := collection.InsertOne(context.Background(), value)

	if insertErr != nil {
		return nil, rest_err.NewInternalServerError(insertErr.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return *converter.ConverterEntityToDomain(*value), nil
}
