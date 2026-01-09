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
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"go.uber.org/zap"
)

type HerdService struct {
	repo     *repository.TransactionalRepository
	services *Service
}

func (s *HerdService) GetHorses(ctx context.Context, herdID uuid.UUID, limit int32, page int32, search string) (horses []*models.Horse, totalCount int32, err error) {
	exec := s.repo.GetExecutor(ctx)

	// Используем срез модификаторов для динамического построения запроса [1]
	mods := []bob.Mod[*dialect.SelectQuery]{
		sm.Where(schema.Horses.Columns.Herd.EQ(psql.S(herdID.String()))),
	}

	// Динамическое добавление фильтра поиска [2, 3]
	if search != "" {
		searchPattern := psql.S("%" + search + "%")
		mods = append(mods, sm.Where(
			schema.Horses.Columns.Name.ILike(searchPattern),
		))
	}

	// 1. Получаем общее количество подходящих записей через Count() [5, 6]
	count, err := schema.Horses.Query(mods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}
	totalCount = int32(count)

	// 2. Добавляем пагинацию к существующим фильтрам [1]
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	paginationMods := append(mods,
		sm.Limit(int(limit)),
		sm.Offset(int(offset)),
	)

	// 3. Выполняем итоговый запрос [5, 6]
	dbHorses, err := schema.Horses.Query(paginationMods...).All(ctx, exec)
	if err != nil {
		zap.L().Error("failed to fetch horses", zap.Error(err))
		return nil, 0, err
	}

	horses = make([]*models.Horse, len(dbHorses))
	for i, horse := range dbHorses {
		horses[i] = &models.Horse{}
		horses[i].Load(*horse)
	}

	return horses, totalCount, nil
}

func (s *HerdService) Create(ctx context.Context, name string, description *string, accountID uuid.UUID) (*models.Herd, error) {
	exec := s.repo.GetExecutor(ctx)

	dbHerd, err := schema.Herds.Insert(&schema.HerdSetter{
		ID:          omit.From(uuid.New()),
		Name:        omit.From(name),
		Description: omitnull.FromPtr(description),
		AccountID:   omit.From(accountID),
	}).One(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	herd := models.Herd{}

	herd.LoadFromDB(*dbHerd)

	return &herd, nil
}

func (s *HerdService) GetAll(ctx context.Context, accountID uuid.UUID, limit int32, page int32, search string) (herds []*models.Herd, totalCount int32, err error) {
	exec := s.repo.GetExecutor(ctx)

	var mods []bob.Mod[*dialect.SelectQuery]

	if search != "" {
		searchPattern := psql.S("%" + search + "%")
		mods = append(mods, sm.Where(
			schema.Herds.Columns.Name.ILike(searchPattern),
		))
	}

	mods = append(mods, sm.Where(
		schema.Herds.Columns.AccountID.EQ(psql.S(accountID.String())),
	))

	count, err := schema.Herds.Query(mods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}
	totalCount = int32(count)

	offset := (page - 1) * limit
	paginationMods := append(mods,
		sm.Limit(int(limit)),
		sm.Offset(int(offset)),
	)

	dbHerds, err := schema.Herds.Query(paginationMods...).All(ctx, exec)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, 0, err
	}

	herds = make([]*models.Herd, len(dbHerds))
	for i, dbHerd := range dbHerds {
		herds[i] = &models.Herd{}
		herds[i].LoadFromDB(*dbHerd)
	}

	return herds, totalCount, nil
}

func (s *HerdService) GetByID(ctx context.Context, id uuid.UUID) (*models.Herd, error) {
	exec := s.repo.GetExecutor(ctx)

	dbHerd, err := schema.FindHerd(ctx, exec, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		zap.L().Error(err.Error())
		return nil, err
	}

	herd := &models.Herd{}
	herd.LoadFromDB(*dbHerd)

	return herd, nil
}

func NewHerdService(repo *repository.TransactionalRepository, services *Service) *HerdService {
	return &HerdService{repo: repo, services: services}
}
