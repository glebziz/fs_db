package db

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	store "github.com/glebziz/fs_db/internal/proto"
)

type db struct {
	client store.StoreV1Client
}

func New(ctx context.Context, url string) (*db, error) {
	conn, err := grpc.DialContext(
		ctx, fmt.Sprintf(url),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}

	return &db{
		client: store.NewStoreV1Client(conn),
	}, nil
}
