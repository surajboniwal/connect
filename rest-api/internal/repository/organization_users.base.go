package repository

import (
	"connect-rest-api/internal/util/apperror"
)

type OrganizationUsersRepository interface {
	AddUser(int64, int64) *apperror.AppError
}
