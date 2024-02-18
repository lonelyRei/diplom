package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/entities"
	"net/http"
)

func (h *Handler) register(context *gin.Context) {
	var user entities.User
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) login(context *gin.Context) {
	var input entities.AuthUserInput

	if err := context.BindJSON(&input); err != nil {
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password)

	if err != nil {
		sendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *Handler) test(context *gin.Context) {
	var test entities.Test
	if err := context.BindJSON(&test); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.services.Test(test)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userCtx, err := getUserContext(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"claims": userCtx.UserId})
}
