package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/intwone/golang-api/src/configuration/database/mongodb"
	"github.com/intwone/golang-api/src/configuration/logger"
	"github.com/intwone/golang-api/src/controller/routes"
	"github.com/intwone/golang-api/src/util"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())

	if err != nil {
		message := fmt.Sprintf("error to connect with database = %s", err.Error())
		logger.Error(message, err, util.CreateJourneyField("main"))
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
