package service

import (
	"context"
	"horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"go.uber.org/zap"
)

type GeneticMarkerService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *GeneticMarkerService) GetAll(ctx context.Context) ([]*models.GeneticMarker, error) {
	exec := s.repo.GetExecutor(ctx)

	dbGeneticMarkers, err := schema.GeneticMarkers.Query().All(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	markers := make([]*models.GeneticMarker, len(dbGeneticMarkers))
	for i, marker := range dbGeneticMarkers {
		markers[i] = &models.GeneticMarker{}
		markers[i].LoadFromDB(*marker)
	}

	return markers, nil
}

func NewGeneticMarkerService(repo *repository.TransactionalRepository, services *Service) *GeneticMarkerService {
	return &GeneticMarkerService{repo: repo, services: services}
}
