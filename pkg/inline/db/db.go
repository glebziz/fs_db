package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/config"
	dbManager "github.com/glebziz/fs_db/internal/db"
	"github.com/glebziz/fs_db/internal/model"
	contentRepo "github.com/glebziz/fs_db/internal/repository/content"
	contentFileRepo "github.com/glebziz/fs_db/internal/repository/content_file"
	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
	fileRepo "github.com/glebziz/fs_db/internal/repository/file"
	txRepo "github.com/glebziz/fs_db/internal/repository/transaction"
	cleanerUseCase "github.com/glebziz/fs_db/internal/usecase/cleaner"
	dirUseCase "github.com/glebziz/fs_db/internal/usecase/dir"
	rootUseCase "github.com/glebziz/fs_db/internal/usecase/root"
	storeUseCase "github.com/glebziz/fs_db/internal/usecase/store"
	txUseCase "github.com/glebziz/fs_db/internal/usecase/transaction"
	"github.com/glebziz/fs_db/internal/utils/disk"
	"github.com/glebziz/fs_db/internal/utils/generator"
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true

type storeUsecase interface {
	Set(ctx context.Context, key string, content *model.Content) error
	Get(ctx context.Context, key string) (*model.Content, error)
	Delete(ctx context.Context, key string) error
}

type txUsecase interface {
	Begin(ctx context.Context, isoLevel model.TxIsoLevel) (string, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type db struct {
	sUc  storeUsecase
	txUc txUsecase

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
	contentFileRep := contentFileRepo.New(manager)
	dirRep := dirRepo.New(manager)
	fileRep := fileRepo.New(manager)
	txRep := txRepo.New()

	cleaner := cleanerUseCase.New(contentRep, contentFileRep)

	rootUc := rootUseCase.New(cfg.RootDirs, disk.GetDisk(), dirRep)
	dirUc := dirUseCase.New(cfg.MaxDirCount, rootUc, dirRep, gen)
	storeUc := storeUseCase.New(
		dirUc, contentRep,
		contentFileRep, fileRep,
		txRep, gen,
	)
	txUx := txUseCase.New(cleaner, fileRep, txRep, gen)

	return &db{
		sUc:     storeUc,
		txUc:    txUx,
		manager: manager,
	}, nil
}

func (db *db) GetStoreUseCase() storeUsecase {
	return db.sUc
}

func (db *db) GetTxUseCase() txUsecase {
	return db.txUc
}

func (db *db) Close() error {
	return db.manager.Close()
}
