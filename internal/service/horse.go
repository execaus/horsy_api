package service

import (
	"context"
	"errors"
	schema "horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type CreateHorseParams struct {
	Herd           uuid.UUID
	Gender         *uuid.UUID
	Name           *string
	BirthDay       *int32
	BirthMonth     *int32
	BirthYear      *int32
	BirthPlace     *uuid.UUID
	WithersHeight  *int32
	Sire           *uuid.UUID
	Dam            *uuid.UUID
	IsPregnant     bool
	Description    *string
	Breeds         []CreateHorseBreedParams
	Color          uuid.UUID
	GeneticMarkers []uuid.UUID
}

type HorseService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *HorseService) GetHorseByID(ctx context.Context, id uuid.UUID) (*models.Horse, error) {
	return s.getHorseByID(ctx, id)
}

func (s *HorseService) GetRelativeHorses(ctx context.Context, id uuid.UUID) ([]*models.Horse, error) {
	exec := s.repo.GetExecutor(ctx)

	dbHorses, err := repository.GetRelativeHorses(id).All(ctx, exec)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	horses := make([]*models.Horse, len(dbHorses))
	for i, horse := range dbHorses {
		h, err := s.getHorseByID(ctx, horse.ID)
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}

		horses[i] = h
	}

	return horses, nil
}

func (s *HorseService) Create(ctx context.Context, params CreateHorseParams) (*models.Horse, error) {
	exec := s.repo.GetExecutor(ctx)

	dbHorse, err := schema.Horses.Insert(&schema.HorseSetter{
		ID:            omit.From(uuid.New()),
		Herd:          omit.From(params.Herd),
		Gender:        omitnull.FromPtr(params.Gender),
		Name:          omitnull.FromPtr(params.Name),
		BirthDay:      omitnull.FromPtr(params.BirthDay),
		BirthMonth:    omitnull.FromPtr(params.BirthMonth),
		BirthYear:     omitnull.FromPtr(params.BirthYear),
		BirthPlace:    omitnull.FromPtr(params.BirthPlace),
		WithersHeight: omitnull.FromPtr(params.WithersHeight),
		Sire:          omitnull.FromPtr(params.Sire),
		Dam:           omitnull.FromPtr(params.Dam),
		IsPregnant:    omit.From(params.IsPregnant),
		IsDead:        omit.From(false),
		Description:   omitnull.FromPtr(params.Description),
	}).One(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	for _, breed := range params.Breeds {
		_, err = schema.HorseBreeds.Insert(&schema.HorseBreedSetter{
			Horse:   omit.From(dbHorse.ID),
			Breed:   omit.From(breed.ID),
			Percent: omit.From(breed.Percent),
		}).Exec(ctx, exec)
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	_, err = schema.HorseColors.Insert(&schema.HorseColorSetter{
		Horse: omit.From(dbHorse.ID),
		Color: omit.From(params.Color),
	}).Exec(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	for _, marker := range params.GeneticMarkers {
		_, err = schema.HorseGeneticMarkers.Insert(&schema.HorseGeneticMarkerSetter{
			Horse:  omit.From(dbHorse.ID),
			Marker: omit.From(marker),
		}).Exec(ctx, exec)
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	horse, err := s.getHorseByID(ctx, dbHorse.ID)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	return horse, nil
}

func (s *HorseService) getHorseByID(ctx context.Context, id uuid.UUID) (*models.Horse, error) {
	exec := s.repo.GetExecutor(ctx)

	horse := &models.Horse{}

	dbHorse, err := schema.FindHorse(ctx, exec, id)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	horse.Load(*dbHorse)

	herd, err := s.services.Herd.GetByID(ctx, dbHorse.Herd)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	horse.Herd = *herd

	if dbHorse.Gender.IsValue() {
		gender, err := s.services.Gender.GetByID(ctx, dbHorse.Gender.GetOrZero())
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
		horse.Gender = gender
	}

	if dbHorse.BirthPlace.IsValue() {
		birthplace, err := s.services.Birthplace.GetByID(ctx, dbHorse.BirthPlace.GetOrZero())
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
		horse.BirthPlace = birthplace
	}

	if dbHorse.Sire.IsValue() {
		sire := &models.Horse{}
		dbSire, err := schema.FindHorse(ctx, exec, dbHorse.Sire.GetOrZero())
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
		sire.Load(*dbSire)
		horse.Sire = sire
	}

	if dbHorse.Dam.IsValue() {
		dam := &models.Horse{}
		dbDam, err := schema.FindHorse(ctx, exec, dbHorse.Dam.GetOrZero())
		if err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
		dam.Load(*dbDam)
		horse.Dam = dam
	}

	return horse, nil
}

func NewHorseService(repo *repository.TransactionalRepository, services *Service) *HorseService {
	return &HorseService{repo: repo, services: services}
}
