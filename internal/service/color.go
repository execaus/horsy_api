package service

import (
	"context"
	"horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"go.uber.org/zap"
)

type ColorService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *ColorService) GetAll(ctx context.Context) ([]*models.Color, error) {
	exec := s.repo.GetExecutor(ctx)

	dbColors, err := schema.Colors.Query().All(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	colors := make([]*models.Color, len(dbColors))

	for i, color := range dbColors {
		colors[i] = &models.Color{}
		colors[i].LoadFromDB(*color)
	}

	return colors, nil
}

func NewColorService(repo *repository.TransactionalRepository, services *Service) *ColorService {
	return &ColorService{repo: repo, services: services}
}
