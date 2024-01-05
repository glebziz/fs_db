package db

import (
	"context"
	"fmt"
	"io"
)

func (db *db) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	content, err := db.sUc.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("store usecase get: %w", err)
	}

	return content.Reader, nil
}
