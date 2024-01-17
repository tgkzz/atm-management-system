package repository

import (
	"atm/internal/repository/auth"
	"database/sql"
)

type Repository struct {
	auth.IAuthRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IAuthRepo: auth.NewAuthRepo(db),
	}
}
