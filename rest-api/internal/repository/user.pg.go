package repository

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/util/apperror"
	"connect-rest-api/internal/util/idgen"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

func (r UserRepositoryPg) Create(user model.User) (*model.User, *apperror.AppError) {
	user.Id = r.idgen.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	if err != nil {
		return nil, apperror.ParseError(err)
	}

	_, err = r.db.Exec("INSERT INTO users (id, name, email, password, phone) VALUES($1, $2, $3, $4, $5)", user.Id, user.Name, user.Email, user.Password, user.Phone)

	if err != nil {
		return nil, apperror.ParseError(err)
	}

	return &user, nil
}
