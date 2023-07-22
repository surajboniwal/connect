package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
)

type UserRepository interface {
	Create(model.User) (*model.User, *apperror.AppError)
}
