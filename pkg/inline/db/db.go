package db

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/reugn/go-quartz/quartz"

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
)

//go:generate mockgen -source service.go -destination mocks/mocks.go -typed true

type cleaner interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
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
	cleaner cleaner

	sUc  storeUsecase
	txUc txUsecase

	manager io.Closer
}

func New(ctx context.Context, cfg *config.Storage) (*db, error) {
	if err := cfg.Valid(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	manager, err := badger.New(cfg.DbPath)
	if err != nil {
		return nil, fmt.Errorf("db new: %w", err)
	}

	gen := generator.New()

	sched := quartz.NewStdScheduler()

	contentRep := contentRepo.New()
	contentFileRep := contentFileRepo.New(manager)
	fileRep := fileRepo.New(manager)
	txRep := txRepo.New()

	dirRep, err := dirRepo.New(cfg.RootDirs)
	if err != nil {
		return nil, fmt.Errorf("dir new: %w", err)
	}

	coreUseCase := core.New(fileRep)

	cl := cleanerUseCase.New(
		sched, contentRep,
		contentFileRep, coreUseCase,
		txRep,
	)

	dirUc := dirUseCase.New(cfg.MaxDirCount, dirRep, gen)
	storeUc := storeUseCase.New(
		dirUc, contentRep,
		contentFileRep, coreUseCase,
		txRep, gen,
		rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	)
	txUx := txUseCase.New(coreUseCase, txRep, gen)

	err = cl.Run(ctx)
	if err != nil {
		return nil, fmt.Errorf("cleaner run: %w", err)
	}

	return &db{
		cleaner: cl,
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
	err := db.cleaner.Stop(context.Background())
	if err != nil {
		return fmt.Errorf("cleaner stop: %w", err)
	}

	err = db.manager.Close()
	if err != nil {
		return fmt.Errorf("manager close: %w", err)
	}

	return nil
}
