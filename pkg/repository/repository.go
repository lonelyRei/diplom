package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/diplomApi/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	Test(test entities.Test) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
