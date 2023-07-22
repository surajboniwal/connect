package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	if err != nil {
		return nil, util.ParseError(err)
	}

	_, err = r.DB.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, util.ParseError(err)
	}

	return &user, nil
}
