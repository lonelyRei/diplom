package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/entities"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

type authUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(context *gin.Context) {
	go func() {
		var user authUserInput
		if err := context.BindJSON(&user); err != nil {
			sendErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}
	}()
}

func (h *Handler) test(context *gin.Context) {
	var test entities.Test
	if err := context.BindJSON(&test); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Test(test)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"id": id})
}
