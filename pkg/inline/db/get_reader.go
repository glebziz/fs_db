package db

import (
	"context"
	"io"
)

func (db *db) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	return db.Open(ctx, key)
}
