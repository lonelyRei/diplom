package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lonelyRei/diplomApi/pkg/repository"
)

func initPostgresDB() (*sqlx.DB, error) {
	// todo пока что хардкодом, в будущем поправить
	return repository.NewPostgresDB(repository.DBConfig{
		Host:     "localhost",
		Username: "postgres",
		Port:     "5436",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
}
