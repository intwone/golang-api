package mongodb

import (
	"context"
	"fmt"
	"os"

	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL           = "MONGODB_URL"
	MONGODB_DATABASE_NAME = "MONGODB_DATABASE_NAME"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(MONGODB_URL)))

	if err != nil {
		message := fmt.Sprintf("error to connect in mongodb, error = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("MongoDbConnection"))
		return nil, err
	}

	clientPingErr := client.Ping(ctx, nil)

	if clientPingErr != nil {
		message := fmt.Sprintf("error to ping client, error = %s", clientPingErr.Error())
		logger.Error(message, clientPingErr, util.CreateJourneyField("MongoDbConnection"))
		return nil, clientPingErr
	}

	logger.Info("mongodb connected")

	return client.Database(os.Getenv(MONGODB_DATABASE_NAME)), nil

}
