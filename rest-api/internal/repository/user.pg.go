package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
	"connect-rest-api/internal/util/idgen"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryPg struct {
	db    *sqlx.DB
	idgen idgen.IdGen
}

func NewUserRepositoryPg(DB *sqlx.DB, idgen idgen.IdGen) UserRepositoryPg {
	return UserRepositoryPg{
		db:    DB,
		idgen: idgen,
	}
}

func (r UserRepositoryPg) Create(user *model.User) *apperror.AppError {
	user.Id = r.idgen.New()

	err := user.HashPassword()

	if err != nil {
		return apperror.Parse(err)
	}

	_, err = r.db.Exec("INSERT INTO users (id, name, email, password, phone) VALUES($1, $2, $3, $4, $5)", user.Id, user.Name, user.Email, user.Password, user.Phone)

	if err != nil {
		return apperror.Parse(err)
	}

	return nil
}

func (r UserRepositoryPg) GetByEmail(email string) (*model.User, *apperror.AppError) {

	var user model.User

	if err := r.db.QueryRowx("SELECT * from users WHERE email=$1;", email).StructScan(&user); err != nil {
		return nil, apperror.Parse(err)
	}

	return &user, nil
}
