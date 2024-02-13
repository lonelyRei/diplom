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
	h.requestChannel <- Request{Method: "register", Data: user, Context: context}
}

type authUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(context *gin.Context) {
	var user authUserInput
	if err := context.BindJSON(&user); err != nil {
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	h.requestChannel <- Request{Method: "login", Data: user, Context: context}
}
