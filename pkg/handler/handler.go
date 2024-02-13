package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/entities"
	"github.com/lonelyRei/diplomApi/pkg/service"
	"net/http"
)

type Request struct {
	Method  string
	Data    interface{}
	Context *gin.Context
}

type Handler struct {
	services       *service.Service
	requestChannel chan Request
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services:       services,
		requestChannel: make(chan Request),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	h.groupAuthRoutes(router)
	return router
}

func (h *Handler) StartWorkerPool(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go h.worker(i)
	}
}

func (h *Handler) worker(id int) {
	for request := range h.requestChannel {
		switch request.Method {
		case "register":
			user := request.Data.(entities.User)
			_, err := h.services.Authorization.CreateUser(user)
			if err != nil {
				sendErrorResponse(request.Context, http.StatusInternalServerError, err.Error())
				continue
			}
			// Отправка ответа на запрос
			request.Context.JSON(http.StatusOK, id)
		case "login":
			// Обработка запроса авторизации
			// ...
			// Другие обработчики
		}
	}
}

func (h *Handler) groupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}
}
