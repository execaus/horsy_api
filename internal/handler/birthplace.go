package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getBirthplaces(c *gin.Context) {
	var (
		err         error
		birthplaces []*models.Birthplace
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		birthplaces, err = services.Birthplace.GetAll(ctx)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	outputBirthplaces := make([]*dto.BirthplaceOutput, len(birthplaces))
	for i, birthplace := range birthplaces {
		outputBirthplaces[i] = &dto.BirthplaceOutput{}
		outputBirthplaces[i].Load(*birthplace)
	}

	sendOK(c, &dto.GetBirthplacesResponse{Birthplaces: outputBirthplaces})
}
