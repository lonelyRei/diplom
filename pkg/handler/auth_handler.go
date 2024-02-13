package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/entities"
	"net/http"
)

func (h *Handler) register(context *gin.Context) {
	var user entities.User

	if err := context.BindJSON(&user); err != nil {
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(user)

	if err != nil {
		sendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type authUserInput struct {
	Username string `json:"username" binding:"required"`

	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(context *gin.Context) {
	var user authUserInput

	if err := context.BindJSON(&user); err != nil {
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	// todo подумать о необходимости двух токенов
	context.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
