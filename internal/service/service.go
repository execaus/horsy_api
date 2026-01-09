package service

import (
	"context"
	"horsy_api/config"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"github.com/google/uuid"
)

type Herd interface {
	Create(ctx context.Context, name string, description *string, accountID uuid.UUID) (*models.Herd, error)
	GetAll(ctx context.Context, accountID uuid.UUID, limit int32, page int32, search string) (herds []*models.Herd, totalCount int32, err error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Herd, error)
	GetHorses(ctx context.Context, id uuid.UUID, limit int32, page int32, search string) (horses []*models.Horse, totalCount int32, err error)
}

type Gender interface {
	Create(ctx context.Context, name string, description *string) (*models.Gender, error)
	GetAll(ctx context.Context) ([]*models.Gender, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Gender, error)
}

type Color interface {
	GetAll(ctx context.Context) ([]*models.Color, error)
}

type Birthplace interface {
	GetAll(ctx context.Context) ([]*models.Birthplace, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Birthplace, error)
}

type GeneticMarker interface {
	GetAll(ctx context.Context) ([]*models.GeneticMarker, error)
}

type Breed interface {
	GetAll(ctx context.Context) ([]*models.Breed, error)
}

type Email interface {
	SendCreatedAccountMail(ctx context.Context, email, password string) error
}

type Account interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Account, error)
	GetByEmail(ctx context.Context, email string) (*models.Account, error)
	Create(ctx context.Context, email string) (token string, err error)
	GetToken(ctx context.Context, email, password string) (token string, err error)
}

type Auth interface {
	GenerateToken(ctx context.Context, accountID uuid.UUID) (string, error)
	GetClaimsFromToken(ctx context.Context, token string) (*models.AuthClaims, error)
	ComparePassword(ctx context.Context, hashedPassword string, password string) bool
	hashPassword(ctx context.Context, password string) (string, error)
	generatePassword() (string, error)
}

type Horse interface {
	Create(ctx context.Context, horse CreateHorseParams) (*models.Horse, error)
	GetHorseByID(ctx context.Context, id uuid.UUID) (*models.Horse, error)
	GetRelativeHorses(ctx context.Context, id uuid.UUID) ([]*models.Horse, error)
}

type Service struct {
	Account
	Auth
	Herd
	Gender
	Color
	Birthplace
	GeneticMarker
	Breed
	Email
	Horse
	repo repository.Transactable
}

func NewService(cfg config.Config, r *repository.TransactionalRepository) *Service {
	s := &Service{}

	s.repo = r
	s.Account = NewAccountService(r, s)
	s.Auth = NewAuthService(cfg.Auth)
	s.Herd = NewHerdService(r, s)
	s.Gender = NewGenderService(r, s)
	s.Color = NewColorService(r, s)
	s.Birthplace = NewBirthplaceService(r, s)
	s.GeneticMarker = NewGeneticMarkerService(r, s)
	s.Breed = NewBreedService(r, s)
	s.Email = NewEmailService(cfg.Email)
	s.Horse = NewHorseService(r, s)

	return s
}
