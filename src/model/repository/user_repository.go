package repository

import (
	"context"

	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.database.Collection("users")

	value, stringfyErr := userDomain.GetJSONValue()

	if stringfyErr != nil {
		return nil, rest_err.NewInternalServerError(stringfyErr.Error())
	}

	result, insertErr := collection.InsertOne(context.Background(), value)

	if insertErr != nil {
		return nil, rest_err.NewInternalServerError(insertErr.Error())
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil
}
