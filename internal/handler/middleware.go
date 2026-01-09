package handler

import (
	"context"
	"horsy_api/internal/models"
	"horsy_api/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type contextKey string

const AccountIDKey contextKey = "accountID"

func (h *Handler) authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		zap.L().Error("missing or invalid Authorization header")
		sendUnauthorized(c)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var (
		err    error
		claims *models.AuthClaims
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		claims, err = services.Auth.GetClaimsFromToken(ctx, token)
		return err
	}); err != nil {
		zap.L().Error(err.Error())
		sendUnauthorized(c)
		return
	}

	if claims == nil {
		sendUnauthorized(c)
		return
	}

	c.Set(AccountIDKey, claims.AccountID)
	c.Next()
}

func getAccountIDFromContext(ctx *gin.Context) (uuid.UUID, bool) {
	accountID, ok := ctx.Get(AccountIDKey)
	if !ok {
		return uuid.Nil, false
	}
	idStr, ok := accountID.(uuid.UUID)
	return idStr, ok
}
