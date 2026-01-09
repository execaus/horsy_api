package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getBreeds(c *gin.Context) {
	var (
		err    error
		breeds []*models.Breed
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		breeds, err = services.Breed.GetAll(ctx)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	outputBreeds := make([]*dto.BreedOutput, len(breeds))
	for i, breed := range breeds {
		outputBreeds[i] = &dto.BreedOutput{
			ID:          breed.ID,
			Name:        breed.Name,
			Description: breed.Description,
		}
	}

	sendOK(c, &dto.GetBreedsResponse{Breeds: outputBreeds})
}
