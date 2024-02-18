package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	h.groupAuthRoutes(router)
	return router
}

func (h *Handler) groupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.POST("/test", h.test)
	}
}
