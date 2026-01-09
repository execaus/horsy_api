package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getGeneticMarkers(c *gin.Context) {
	var (
		err     error
		markers []*models.GeneticMarker
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		markers, err = services.GeneticMarker.GetAll(ctx)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	markersOutput := make([]*dto.GeneticMarkerOutput, len(markers))
	for i, marker := range markers {
		markersOutput[i] = &dto.GeneticMarkerOutput{
			ID:          marker.ID,
			Name:        marker.Name,
			Description: marker.Description,
		}
	}

	sendOK(c, &dto.GetGeneticMarkersResponse{GeneticMarkers: markersOutput})
}
