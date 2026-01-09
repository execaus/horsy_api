package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createHorse(c *gin.Context) {
	var in dto.CreateHorseRequest

	if err := c.BindJSON(&in); err != nil {
		sendBadRequest(c, err)
		return
	}

	breeds := make([]service.CreateHorseBreedParams, len(in.Breeds))
	for i, breed := range in.Breeds {
		breeds[i] = service.CreateHorseBreedParams{
			ID:      breed.ID,
			Percent: breed.Percent,
		}
	}

	var (
		err   error
		horse *models.Horse
	)

	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		horse, err = services.Horse.Create(ctx, service.CreateHorseParams{
			Herd:           in.Herd,
			Gender:         in.Gender,
			Name:           in.Name,
			BirthDay:       in.BirthDay,
			BirthMonth:     in.BirthMonth,
			BirthYear:      in.BirthYear,
			BirthPlace:     in.BirthPlace,
			WithersHeight:  in.WithersHeight,
			Sire:           in.Sire,
			Dam:            in.Dam,
			IsPregnant:     in.IsPregnant,
			Description:    in.Description,
			Breeds:         breeds,
			Color:          in.Color,
			GeneticMarkers: in.GeneticMarkers,
		})
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	horseOutput := dto.HorseOutput{}
	horseOutput.Load(horse)

	sendOK(c, &dto.CreateHorseResponse{
		Horse: horseOutput,
	})
}

func (h *Handler) getHorse(c *gin.Context) {
	id, err := h.getParameterUUID(c, "id")
	if err != nil {
		sendBadRequest(c, err)
		return
	}

	var (
		horse          *models.Horse
		relativeHorses []*models.Horse
	)
	err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		horse, err = services.Horse.GetHorseByID(ctx, id)
		if err != nil {
			return err
		}
		relativeHorses, err = services.Horse.GetRelativeHorses(ctx, id)
		return err
	})
	if err != nil {
		sendServiceError(c, err)
		return
	}

	horseOutput := dto.HorseOutput{}
	horseOutput.Load(horse)

	relativeHorsesOutput := make([]dto.HorseOutput, len(relativeHorses))
	for i, relativeHorse := range relativeHorses {
		relativeHorsesOutput[i] = dto.HorseOutput{}
		relativeHorsesOutput[i].Load(relativeHorse)
	}

	sendOK(c, dto.GetHorseResponse{
		Horse:          horseOutput,
		RelativeHorses: relativeHorsesOutput,
	})
}
