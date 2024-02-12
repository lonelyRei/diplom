package handler

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/pkg/service"
	"net/http"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	h.setCors(router)

	h.groupAuthRoutes(router)

	return router
}

func (h *Handler) groupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		// Обработчик регистрации нового пользователя
		auth.POST("/register", func(context *gin.Context) {
			context.JSON(http.StatusOK, map[string]interface{}{"success": true})
		})

		// Обработчик авторизации
		//auth.POST("/login", h.signIn)
	}
}

func (h *Handler) setCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,

		AllowOriginFunc: func(origin string) bool {
			fmt.Println(origin == "http://localhost:5173" || origin == "http://localhost:5173/")
			return origin == "http://localhost:5173" || origin == "http://localhost:5173/"
		},
		MaxAge: 12 * time.Hour,
	}))
}
