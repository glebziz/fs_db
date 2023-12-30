package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/config"
	dbManager "github.com/glebziz/fs_db/internal/db"
	"github.com/glebziz/fs_db/internal/model"
	contentRepo "github.com/glebziz/fs_db/internal/repository/content"
	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
	fileRepo "github.com/glebziz/fs_db/internal/repository/file"
	dirUseCase "github.com/glebziz/fs_db/internal/usecase/dir"
	rootUseCase "github.com/glebziz/fs_db/internal/usecase/root"
	storeUseCase "github.com/glebziz/fs_db/internal/usecase/store"
	"github.com/glebziz/fs_db/internal/utils/disk"
	generator "github.com/glebziz/fs_db/internal/utils/generator"
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true

type useCase interface {
	Set(ctx context.Context, key string, content *model.Content) error
	Get(ctx context.Context, key string) (*model.Content, error)
	Delete(ctx context.Context, key string) error
}

type db struct {
	usecase useCase
	manager *dbManager.Manager
}

func New(ctx context.Context, cfg *config.Storage) (*db, error) {
	if err := cfg.Valid(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	manager, err := dbManager.New(ctx, cfg.DbPath)
	if err != nil {
		return nil, fmt.Errorf("db new: %w", err)
	}

	gen := generator.New()

	contentRep := contentRepo.New()
	dirRep := dirRepo.New(manager)
	fileRep := fileRepo.New(manager)

	rootUc := rootUseCase.New(cfg.RootDirs, disk.GetDisk(), dirRep)
	dirUc := dirUseCase.New(cfg.MaxDirCount, rootUc, dirRep, gen)
	storeUc := storeUseCase.New(dirUc, contentRep, fileRep, gen)

	return &db{
		usecase: storeUc,
		manager: manager,
	}, nil
}

func (db *db) GetUseCase() useCase {
	return db.usecase
}

func (db *db) Close() error {
	return db.manager.Close()
}
