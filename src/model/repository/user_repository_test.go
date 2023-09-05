package repository

import (
	"fmt"
	"testing"

	"github.com/intwone/golang-api/src/model"
	"github.com/intwone/golang-api/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database := "user_database_test"

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("should_return_a_valid_domain_successfully", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUser(model.NewUserDomain("test@mail.com", "1234567", "test name", 20))

		_, errId := primitive.ObjectIDFromHex(userDomain.GetId())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), "test@mail.com")
	})

	mtestDb.Run("should_return_an_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUser(model.NewUserDomain("test@mail.com", "1234567", "test name", 20))

		assert.Nil(t, userDomain)
		assert.NotNil(t, err)
	})
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database := "user_database_test"

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("should_return_success_when_valid_email_provided", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			Id:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "1234567",
			Name:     "test name",
			Age:      20,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", database, "users"), mtest.FirstBatch, convertEntityToBson(userEntity)))

		databaseMock := mt.Client.Database(database)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.NotNil(t, userDomain)
	})

	mtestDb.Run("should_return_and_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("error@mail.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("should_return_an_error_when_user_not_exists", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			Id:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "1234567",
			Name:     "test name",
			Age:      20,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", database, "users"), mtest.FirstBatch))

		databaseMock := mt.Client.Database(database)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.Id},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
