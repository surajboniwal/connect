package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util"
)

type UserRepository interface {
	Create(model.User) (*model.User, *util.AppError)
}
