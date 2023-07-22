package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrganizationRepositoryMongo struct {
	DB *mongo.Collection
}

func NewOrganizationRepositoryMongo(db *mongo.Database) OrganizationRepositoryMongo {
	return OrganizationRepositoryMongo{
		DB: db.Collection("organizations"),
	}
}

func (r OrganizationRepositoryMongo) Create(organization model.Organization) (*model.Organization, *apperror.AppError) {

	organization.Id = primitive.NewObjectID()

	_, err := r.DB.InsertOne(context.TODO(), organization)

	if err != nil {
		return nil, &apperror.AppError{
			OriginalError: err,
		}
	}

	return &organization, nil
}
