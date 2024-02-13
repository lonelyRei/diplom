package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/diplomApi/entities"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user entities.User) (int, error) {
	// todo: implement
	return 1, nil
}
