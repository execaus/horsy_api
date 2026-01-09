package service

import (
	"context"
	"errors"
	"horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type GenderService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *GenderService) GetByID(ctx context.Context, id uuid.UUID) (*models.Gender, error) {
	exec := s.repo.GetExecutor(ctx)

	dbGender, err := schema.FindGender(ctx, exec, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	gender := &models.Gender{}
	gender.LoadFromDB(*dbGender)

	return gender, nil
}

func (s *GenderService) Create(ctx context.Context, name string, description *string) (*models.Gender, error) {
	exec := s.repo.GetExecutor(ctx)

	dbGender, err := schema.Genders.Insert(&schema.GenderSetter{
		ID:          omit.From(uuid.New()),
		Name:        omit.From(name),
		Description: omitnull.FromPtr(description),
	}).One(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	gender := models.Gender{}
	gender.LoadFromDB(*dbGender)

	return &gender, nil
}

func (s *GenderService) GetAll(ctx context.Context) ([]*models.Gender, error) {
	exec := s.repo.GetExecutor(ctx)

	dbGenders, err := schema.Genders.Query().All(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	genders := make([]*models.Gender, len(dbGenders))

	for i, gender := range dbGenders {
		genders[i] = &models.Gender{}
		genders[i].LoadFromDB(*gender)
	}

	return genders, nil
}

func NewGenderService(repo *repository.TransactionalRepository, services *Service) *GenderService {
	return &GenderService{repo: repo, services: services}
}
