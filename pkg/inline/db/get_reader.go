package db

import (
	"context"
	"fmt"
	"io"
)

func (db *db) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	content, err := db.container.Store().Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("store usecase get: %w", err)
	}

	return content, nil
}
