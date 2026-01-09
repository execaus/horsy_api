package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input dto.SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		sendBadRequest(c, err)
		return
	}

	var (
		err   error
		token string
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		token, err = services.Account.Create(ctx, input.Email)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	sendCreated(c, &dto.SignUpResponse{
		Token: token,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input dto.SignInRequest

	if err := c.BindJSON(&input); err != nil {
		sendBadRequest(c, err)
		return
	}

	var (
		err   error
		token string
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		token, err = services.Account.GetToken(ctx, input.Email, input.Password)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	sendOK(c, &dto.SignInResponse{
		Token: token,
	})
}
