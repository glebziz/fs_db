package app

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/di"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/interceptors/server"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

type app struct {
	cfg       config.Config
	container *di.Container
	server    *grpc.Server
}

func New(ctx context.Context, cfg config.Config) (*app, error) {
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

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			server.LoggingInterceptor,
			server.ContextInterceptor,
		),
		grpc.ChainStreamInterceptor(
			server.StreamLoggingInterceptor,
			server.ContextStreamInterceptor,
		),
	)

	store.RegisterStoreV1Server(s, container.StoreService())

	return &app{
		cfg:       cfg,
		container: container,
		server:    s,
	}, nil
}
