package db

import (
	"context"
	"fmt"
	"io"
)

func (db *db) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	content, err := db.usecase.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("usecase get: %w", err)
	}

	return content.Reader, nil
}
