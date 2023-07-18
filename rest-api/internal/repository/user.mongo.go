package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	DB *mongo.Collection
}

func NewUserRepositoryMongo(DB *mongo.Database) UserRepositoryMongo {
	return UserRepositoryMongo{
		DB: DB.Collection("users"),
	}
}

func (r UserRepositoryMongo) Create(user model.User) (*model.User, *util.AppError) {
	user.Id = primitive.NewObjectID()

	_, err := r.DB.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, util.ParseMongoError(err)
	}

	return &user, nil
}
