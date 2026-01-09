package handler

import (
	"context"
	"horsy_api/internal/dto"
	"horsy_api/internal/models"
	"horsy_api/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createHerd(c *gin.Context) {
	var in dto.CreateHerdRequest

	if err := c.BindJSON(&in); err != nil {
		sendBadRequest(c, err)
		return
	}

	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		sendUnauthorized(c)
		return
	}

	var err error
	var herd *models.Herd
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		herd, err = services.Herd.Create(ctx, in.Name, in.Description, accountID)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	herdOutput := dto.HerdOutput{}
	herdOutput.Load(*herd)

	sendOK(c, &dto.CreateHerdResponse{Herd: herdOutput})
}

func (h *Handler) getHerds(c *gin.Context) {
	var params dto.GetListParams

	params.BindFromContext(c)

	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		sendUnauthorized(c)
		return
	}

	var (
		err   error
		herds []*models.Herd
		total int32
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		herds, total, err = services.Herd.GetAll(ctx, accountID, params.Limit, params.Page, params.Search)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	outputHerds := make([]dto.HerdOutput, len(herds))
	for i, herd := range herds {
		outputHerds[i] = dto.HerdOutput{}
		outputHerds[i].Load(*herd)
	}

	sendOK(c, &dto.GetHerdsResponse{
		Herds:      outputHerds,
		TotalCount: total,
	})
}

func (h *Handler) getHerdByID(c *gin.Context) {
	id, err := h.getParameterUUID(c, "id")
	if err != nil {
		sendBadRequest(c, err)
		return
	}

	var (
		herd *models.Herd
	)
	if err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		herd, err = services.Herd.GetByID(ctx, id)
		return err
	}); err != nil {
		sendServiceError(c, err)
		return
	}

	herdOutput := dto.HerdOutput{}
	herdOutput.Load(*herd)

	sendOK(c, &dto.GetHerdByIDResponse{Herd: herdOutput})
}

func (h *Handler) getHerdHorses(c *gin.Context) {
	var params dto.GetListParams

	params.BindFromContext(c)

	id, err := h.getParameterUUID(c, "id")
	if err != nil {
		sendBadRequest(c, err)
		return
	}

	var (
		horses     []*models.Horse
		totalCount int32
	)
	err = h.saga.Run(c, func(ctx context.Context, services *service.Service) error {
		horses, totalCount, err = services.Herd.GetHorses(ctx, id, params.Limit, params.Page, params.Search)
		return err
	})
	if err != nil {
		sendServiceError(c, err)
		return
	}

	output := &dto.GetHerdHorsesResponse{
		Horses:     make([]dto.HorseOutput, len(horses)),
		TotalCount: totalCount,
	}

	for i, horse := range horses {
		output.Horses[i] = dto.HorseOutput{}
		output.Horses[i].Load(horse)
	}

	sendOK(c, output)
}
