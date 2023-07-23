package repository

import (
	"connect-rest-api/internal/util/idgen"

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
