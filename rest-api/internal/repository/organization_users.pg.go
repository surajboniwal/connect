package repository

import (
	"connect-rest-api/internal/util/apperror"

	"github.com/jmoiron/sqlx"
)

type OrganizationUsersRepositoryPg struct {
	db *sqlx.DB
}

func NewOrganizationUsersRepositoryPg(db *sqlx.DB) OrganizationUsersRepositoryPg {
	return OrganizationUsersRepositoryPg{
		db: db,
	}
}

func (r OrganizationUsersRepositoryPg) AddUser(organizationId int64, userId int64) *apperror.AppError {
	_, err := r.db.Exec("INSERT into organization_users (organization_id, user_id) VALUES($1, $2)", organizationId, userId)

	if err != nil {
		return apperror.Parse(err)
	}

	return nil
}
