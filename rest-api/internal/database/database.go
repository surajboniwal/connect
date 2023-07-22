package database

import (
	"connect-rest-api/internal/config"
	"connect-rest-api/internal/util/applogger"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger applogger.Logger = applogger.New("database")

func Connect(config *config.Config) mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DB_URL))

	if err != nil {
		logger.E(err)
	}

	logger.I("Connected to database")

	return *client.Database("connect")
}
