package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
	"connect-rest-api/internal/util/idgen"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryPg struct {
	db    *sqlx.DB
	idgen idgen.IdGen
}

func NewAuthRepositoryPg(db *sqlx.DB, idgen idgen.IdGen) AuthRepositoryPg {
	return AuthRepositoryPg{
		db:    db,
		idgen: idgen,
	}
}

func (r AuthRepositoryPg) Register(user *model.User, organization *model.Organization) *apperror.AppError {

	user.Id = r.idgen.New()
	organization.Id = r.idgen.New()

	err := user.HashPassword()

	if err != nil {
		return apperror.Parse(err)
	}

	tx := r.db.MustBeginTx(context.TODO(), &sql.TxOptions{ReadOnly: false})

	_, err = tx.Exec("INSERT INTO users (id, name, email, password, phone) VALUES($1, $2, $3, $4, $5)", user.Id, user.Name, user.Email, user.Password, user.Phone)

	if err != nil {
		tx.Rollback()
		return apperror.Parse(err)
	}

	_, err = tx.Exec("INSERT into organizations (id, name) VALUES ($1, $2)", organization.Id, organization.Name)

	if err != nil {
		tx.Rollback()
		return apperror.Parse(err)
	}

	_, err = tx.Exec("INSERT into organization_users (organization_id, user_id) VALUES($1, $2)", organization.Id, user.Id)

	if err != nil {
		tx.Rollback()
		return apperror.Parse(err)
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return apperror.Parse(err)
	}

	return nil
}
