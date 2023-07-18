package database

import (
	"connect-rest-api/internal/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config *config.Config) mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DB_URL))

	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}

	fmt.Println("Connected to database")

	return *client.Database("connect")
}
