package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/di"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

type db struct {
	container *di.Container
}

func New(ctx context.Context, cfg config.Config) (*db, error) {
	if err := cfg.Storage.Valid(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	container := di.New(cfg)

	container.Pool().Run(ctx)
	deleteFiles, err := container.Core().Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("load core: %w", err)
	}

	container.Cleaner().DeleteFilesAsync(ctx, deleteFiles)
	container.Pool().Sched(ctx, wpool.Event{
		Caller: "DeleteOld every minute",
		Fn: func(ctx context.Context) error {
			return container.Cleaner().DeleteOld(ctx)
		},
	}, cfg.Storage.GCPeriod)

	return &db{
		container: container,
	}, nil
}

func (db *db) Close() error {
	db.container.Pool().Stop()
	err := db.container.Badger().Close()
	if err != nil {
		return fmt.Errorf("manager close: %w", err)
	}

	return nil
}
