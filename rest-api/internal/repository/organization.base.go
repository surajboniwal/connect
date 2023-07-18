package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util"
)

type OrganizationRepository interface {
	Create(model.Organization) (*model.Organization, *util.AppError)
}
