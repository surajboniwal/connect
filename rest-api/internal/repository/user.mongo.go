package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryMongo struct {
	db *mongo.Collection
}

func NewUserRepositoryMongo(DB *mongo.Database) UserRepositoryMongo {
	return UserRepositoryMongo{
		db: DB.Collection("users"),
	}
}

func (r UserRepositoryMongo) Create(user model.User) (*model.User, *apperror.AppError) {
	user.Id = primitive.NewObjectID()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	if err != nil {
		return nil, apperror.ParseError(err)
	}

	_, err = r.db.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, apperror.ParseError(err)
	}

	return &user, nil
}
