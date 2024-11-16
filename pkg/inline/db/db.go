package db

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model"
	contentRepo "github.com/glebziz/fs_db/internal/repository/content"
	contentFileRepo "github.com/glebziz/fs_db/internal/repository/content_file"
	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
	fileRepo "github.com/glebziz/fs_db/internal/repository/file"
	txRepo "github.com/glebziz/fs_db/internal/repository/transaction"
	cleanerUseCase "github.com/glebziz/fs_db/internal/usecase/cleaner"
	"github.com/glebziz/fs_db/internal/usecase/core"
	dirUseCase "github.com/glebziz/fs_db/internal/usecase/dir"
	storeUseCase "github.com/glebziz/fs_db/internal/usecase/store"
	txUseCase "github.com/glebziz/fs_db/internal/usecase/transaction"
	"github.com/glebziz/fs_db/internal/utils/generator"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

type pool interface {
	Stop()
}

type storeUsecase interface {
	Set(ctx context.Context, key string, content io.Reader) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}

type txUsecase interface {
	Begin(ctx context.Context, isoLevel model.TxIsoLevel) (string, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type db struct {
	pool pool

	sUc  storeUsecase
	txUc txUsecase

	manager io.Closer
}

func New(ctx context.Context, cfg config.Config) (*db, error) { //nolint:funlen // TODO fix
	if err := cfg.Storage.Valid(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	manager, err := badger.New(cfg.Storage.DbPath)
	if err != nil {
		return nil, fmt.Errorf("db new: %w", err)
	}

	gen := generator.New()

	p := wpool.New(wpool.Options{
		NumWorkers:   cfg.WPool.NumWorkers,
		SendDuration: cfg.WPool.SendDuration,
	})

	contentRep := contentRepo.New()
	contentFileRep := contentFileRepo.New(manager)
	fileRep := fileRepo.New(manager)
	txRep := txRepo.New()

	dirRep, err := dirRepo.New(cfg.Storage.RootDirs)
	if err != nil {
		return nil, fmt.Errorf("dir new: %w", err)
	}

	coreUseCase := core.New(fileRep)

	cleaner := cleanerUseCase.New(
		coreUseCase, contentRep,
		contentFileRep, manager,
		dirRep, fileRep, p, txRep,
	)

	dirUc := dirUseCase.New(cfg.Storage.MaxDirCount, dirRep, gen)
	storeUc := storeUseCase.New(
		dirUc, contentRep,
		contentFileRep, coreUseCase,
		txRep, gen,
		rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	)
	txUx := txUseCase.New(cleaner, coreUseCase, txRep, gen)
	p.Run(ctx)

	deleteFiles, err := coreUseCase.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("load core: %w", err)
	}
	cleaner.DeleteFilesAsync(ctx, deleteFiles)
	p.Sched(ctx, wpool.Event{
		Caller: "DeleteOld every minute",
		Fn: func(ctx context.Context) error {
			return cleaner.DeleteOld(ctx)
		},
	}, cfg.Storage.GCPeriod)

	return &db{
		pool:    p,
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
	db.pool.Stop()
	err := db.manager.Close()
	if err != nil {
		return fmt.Errorf("manager close: %w", err)
	}

	return nil
}
