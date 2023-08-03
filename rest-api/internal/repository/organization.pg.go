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

func (r OrganizationRepositoryPg) Create(organization *model.Organization) *apperror.AppError {

	organization.Id = r.idgen.New()

	_, err := r.db.Exec("INSERT into organizations (id, name) VALUES ($1, $2)", organization.Id, organization.Name)

	if err != nil {
		return &apperror.AppError{
			OriginalError: err,
		}
	}

	return nil
}

func (r OrganizationRepositoryPg) GetOrganizationsUsingUserId(userId int64) (*[]model.Organization, *apperror.AppError) {

	var organizations = make([]model.Organization, 0)

	err := r.db.Select(&organizations, "SELECT o.* FROM organizations o JOIN organization_users ou ON o.id = ou.organization_id WHERE ou.user_id = $1;", userId)

	if err != nil {
		return &organizations, apperror.Parse(err)
	}

	return &organizations, nil

}
