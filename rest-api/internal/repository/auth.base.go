package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
)

type AuthRepository interface {
	Register(user *model.User, organization *model.Organization) *apperror.AppError
}
