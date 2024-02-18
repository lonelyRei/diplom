package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/diplomApi/pkg/service"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	tokenType  = "Bearer"
	userCtx    = "userClaims"
)

// userIdentity - middleware для проверки авторизованного пользователя
func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		sendErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		sendErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != tokenType {
		sendErrorResponse(ctx, http.StatusUnauthorized, "wrong token type")
		return
	}

	userClaims, err := h.services.ParseToken(headerParts[1])

	if err != nil {
		sendErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userClaims)
}

func getUserContext(c *gin.Context) (*service.TokenClaims, error) {
	claims, ok := c.Get(userCtx)

	if !ok {
		return nil, errors.New("user claims not found")
	}

	typedClaims, ok := claims.(*service.TokenClaims)

	if !ok {
		return nil, errors.New("user claims are of invalid type")
	}

	return typedClaims, nil
}
