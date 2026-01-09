package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getHorseGenderList(c *gin.Context) {
	var (
		err     error
		genders []*models.Gender
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		genders, err = services.Gender.GetAll(ctx)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	gendersOutput := make([]*dto.GenderOutput, len(genders))
	for i, gender := range genders {
		gendersOutput[i] = &dto.GenderOutput{}
		gendersOutput[i].Load(*gender)
	}

	sendOK(c, &dto.GetHorseGenderListResponse{Genders: gendersOutput})
}
