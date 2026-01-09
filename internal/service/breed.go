package service

import (
	"context"
	"horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"go.uber.org/zap"
)

type BreedService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *BreedService) GetAll(ctx context.Context) ([]*models.Breed, error) {
	exec := s.repo.GetExecutor(ctx)

	dbBreeds, err := schema.Breeds.Query().All(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	breeds := make([]*models.Breed, len(dbBreeds))

	for i, b := range dbBreeds {
		breeds[i] = &models.Breed{}
		breeds[i].LoadFromDB(*b)
	}

	return breeds, nil
}

func NewBreedService(repo *repository.TransactionalRepository, services *Service) *BreedService {
	return &BreedService{repo: repo, services: services}
}
