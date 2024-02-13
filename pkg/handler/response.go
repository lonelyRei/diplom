package handler

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"message"`
}

func sendErrorResponse(context *gin.Context, statusCode int, message string) {
	context.AbortWithStatusJSON(statusCode, Error{message})
}
