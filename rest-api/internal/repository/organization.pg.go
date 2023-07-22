package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
	"connect-rest-api/internal/util/idgen"

	"github.com/jmoiron/sqlx"
)

type OrganizationRepositoryPg struct {
	db    *sqlx.DB
	idgen idgen.IdGen
}

func NewOrganizationRepositoryPg(db *sqlx.DB, idGen idgen.IdGen) OrganizationRepositoryPg {
	return OrganizationRepositoryPg{
		db:    db,
		idgen: idGen,
	}
}

func (r OrganizationRepositoryPg) Create(organization model.Organization) (*model.Organization, *apperror.AppError) {

	organization.Id = r.idgen.New()

	_, err := r.db.Exec("INSERT into organizations (id, name) VALUES ($1, $2)", organization.Id, organization.Name)

	if err != nil {
		return nil, &apperror.AppError{
			OriginalError: err,
		}
	}

	return &organization, nil
}
