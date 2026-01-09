package service

import (
	"context"
	"errors"
	"horsy_api/internal/gen/dberrors"
	"horsy_api/internal/gen/schema"
	"horsy_api/internal/models"
	"horsy_api/internal/repository"

	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"go.uber.org/zap"
)

type AccountService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *AccountService) GetToken(ctx context.Context, email, password string) (token string, err error) {
	account, err := s.services.Account.GetByEmail(ctx, email)
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	if account == nil {
		return "", ErrSignInInvalidCredentials
	}

	isComparePassword := s.services.Auth.ComparePassword(ctx, account.Password, password)

	if !isComparePassword {
		return "", ErrSignInInvalidCredentials
	}

	token, err = s.services.Auth.GenerateToken(ctx, account.ID)
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	return token, nil
}

func (s *AccountService) GetByID(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	exec := s.repo.GetExecutor(ctx)

	dbAccount, err := schema.FindAccount(ctx, exec, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	account := models.Account{}
	if err = account.LoadFromDB(*dbAccount); err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) GetByEmail(ctx context.Context, email string) (*models.Account, error) {
	exec := s.repo.GetExecutor(ctx)

	dbAccount, err := schema.Accounts.Query(sm.Where(schema.Accounts.Columns.Email.EQ(psql.S(email)))).One(ctx, exec)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	account := models.Account{}
	if err = account.LoadFromDB(*dbAccount); err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) Create(ctx context.Context, email string) (token string, err error) {
	exec := s.repo.GetExecutor(ctx)

	password, err := s.services.Auth.generatePassword()
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	passwordHash, err := s.services.Auth.hashPassword(ctx, password)
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	dbAccount, err := schema.Accounts.Insert(&schema.AccountSetter{
		ID:       omit.From(uuid.New()),
		Email:    omit.From(email),
		Password: omit.From(passwordHash),
	}).One(ctx, exec)
	if err != nil {
		if errors.Is(err, dberrors.AccountErrors.ErrUniqueAccountsEmailKey) {
			zap.L().Error(ErrAccountEmailExists.Error())
			return "", ErrAccountEmailExists
		}
		zap.L().Error(err.Error())
		return "", err
	}

	token, err = s.services.Auth.GenerateToken(ctx, dbAccount.ID)
	if err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	if err = s.sendNewAccountPassword(ctx, email, password); err != nil {
		zap.L().Error(err.Error())
		return "", err
	}

	return token, nil
}

func NewAccountService(repo *repository.TransactionalRepository, services *Service) *AccountService {
	return &AccountService{
		repo:     repo,
		services: services,
	}
}

func (s *AccountService) sendNewAccountPassword(ctx context.Context, email, password string) error {
	var err error

	if err = s.services.Email.SendCreatedAccountMail(ctx, email, password); err == nil {
		return nil
	}
	zap.L().Error(err.Error())

	return err
}
