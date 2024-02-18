package main

import (
	"github.com/lonelyRei/diplomApi/pkg/handler"
	"github.com/lonelyRei/diplomApi/pkg/repository"
	"github.com/lonelyRei/diplomApi/pkg/service"
	"github.com/lonelyRei/diplomApi/server"
	"log"
)

func main() {
	db, err := initPostgresDB()

	if err != nil {
		log.Fatalf("произошла ошибка при подключении к базе данных: %s", err.Error())
	}

	appRepo := repository.NewRepository(db)

	appServices := service.NewService(appRepo)

	appHandler := handler.NewHandler(appServices)

	srv := new(server.Server)

	if err := srv.Run("80", appHandler.InitRoutes()); err != nil {
		log.Fatalf("произошла ошибка при работе сервера: %s", err.Error())
	}
}
