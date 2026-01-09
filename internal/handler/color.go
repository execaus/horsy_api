package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getColors(c *gin.Context) {
	var (
		err    error
		colors []*models.Color
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		colors, err = services.Color.GetAll(ctx)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	outputColors := make([]*dto.ColorOutput, len(colors))
	for i, color := range colors {
		outputColors[i] = &dto.ColorOutput{
			ID:          color.ID,
			Name:        color.Name,
			Description: color.Description,
		}
	}

	sendOK(c, &dto.GetColorsResponse{Colors: outputColors})
}
