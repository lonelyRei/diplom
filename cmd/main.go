package main

import (
	"github.com/lonelyRei/diplomApi/pkg/handler"
	"github.com/lonelyRei/diplomApi/pkg/repository"
	"github.com/lonelyRei/diplomApi/pkg/service"
	"github.com/lonelyRei/diplomApi/server"
	"log"
)

func main() {
	appRepo := repository.NewRepository(nil)

	appServices := service.NewService(appRepo)

	appHandler := handler.NewHandler(appServices)

	srv := new(server.Server)

	appHandler.StartWorkerPool(5)

	// todo Пока что хардкодом, в будущем исправить
	if err := srv.Run("80", appHandler.InitRoutes()); err != nil {
		log.Fatalf("error happened while running server: %s", err.Error())
	}
}
