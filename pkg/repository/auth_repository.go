package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/diplomApi/entities"
)

const usersTableName = "users"
const testTableName = "test"

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user entities.User) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash, image, age, weight, height) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		usersTableName)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, user.Image, user.Age, user.Weight, user.Height)

	if err := row.Scan(&id); err != nil {
		// Если при выполнении запроса возникла ошибка, то возвращаем ее на слой выше
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) Test(test entities.Test) (int, error) {

	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (username) values ($1) RETURNING id",
		testTableName)

	row := r.db.QueryRow(query, test.Username)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
