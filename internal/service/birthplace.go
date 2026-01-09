package service

import (
	"context"
	"errors"
	"horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type BirthplaceService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *BirthplaceService) GetByID(ctx context.Context, id uuid.UUID) (*models.Birthplace, error) {
	exec := s.repo.GetExecutor(ctx)

	dbBirthplace, err := schema.FindBirthplace(ctx, exec, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	birthplace := &models.Birthplace{}
	birthplace.LoadFromDB(*dbBirthplace)

	return birthplace, nil
}

func (s *BirthplaceService) GetAll(ctx context.Context) ([]*models.Birthplace, error) {
	exec := s.repo.GetExecutor(ctx)

	dbBirthplaces, err := schema.Birthplaces.Query().All(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	birthplaces := make([]*models.Birthplace, len(dbBirthplaces))
	for i, b := range dbBirthplaces {
		birthplaces[i] = &models.Birthplace{}
		birthplaces[i].LoadFromDB(*b)
	}

	return birthplaces, nil
}

func NewBirthplaceService(repo *repository.TransactionalRepository, services *Service) *BirthplaceService {
	return &BirthplaceService{repo: repo, services: services}
}
