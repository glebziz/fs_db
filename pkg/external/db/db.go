package db

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	store "github.com/glebziz/fs_db/internal/proto"
)

type db struct {
	client store.StoreV1Client
}

func New(_ context.Context, url string, opt ...OptionFunc) (*db, error) {
	var opts options
	for _, o := range opt {
		o(&opts)
	}

	credentials, err := opts.credentials()
	if err != nil {
		return nil, fmt.Errorf("build credentials: %w", err)
	}

	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(credentials),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}

	return &db{
		client: store.NewStoreV1Client(conn),
	}, nil
}

func (db *db) Close() error {
	return nil
}
