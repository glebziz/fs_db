package model

import (
	"context"
	"io"
)

type DB interface {
	Set(ctx context.Context, key string, b []byte) error
	SetReader(ctx context.Context, key string, reader io.Reader, size uint64) error
	Get(ctx context.Context, key string) ([]byte, error)
	GetReader(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}
