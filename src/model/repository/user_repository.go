package repository

import (
	"context"
	"fmt"

	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/configuration/rest_err"
	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/model/repository/entity"
	"github.com/intwone/golang-api/src/model/repository/entity/converter"
	"github.com/intwone/golang-api/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.database.Collection("users")

	value := converter.ConverterDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	return *converter.ConverterEntityToDomain(*value), nil
}

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.database.Collection("users")

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			message := fmt.Sprintf("User not found with this email %s", email)
			logger.Info(message, util.CreateJourneyField("UserRepository"))
			return nil, rest_err.NewNotFoundError(message)
		}

		message := fmt.Sprintf("error to FindUserByEmail in repository, error = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("UserRepository"))
		return nil, rest_err.NewInternalServerError(message)
	}

	return *converter.ConverterEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.database.Collection("users")

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			message := fmt.Sprintf("User not found with this id %s", id)
			logger.Info(message, util.CreateJourneyField("UserRepository"))
			return nil, rest_err.NewNotFoundError(message)
		}

		message := fmt.Sprintf("error to FindUserById in repository, error = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("UserRepository"))
		return nil, rest_err.NewInternalServerError(message)
	}

	return *converter.ConverterEntityToDomain(*userEntity), nil
}

func (ur *userRepository) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	collection := ur.database.Collection("users")

	value := converter.ConverterDomainToEntity(userDomain)

	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {
	collection := ur.database.Collection("users")

	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
