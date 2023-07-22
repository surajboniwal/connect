package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
)

type OrganizationRepository interface {
	Create(model.Organization) (*model.Organization, *apperror.AppError)
}
