package db

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

func (db *db) Set(ctx context.Context, key string, b []byte) error {
	err := db.usecase.Set(ctx, key, &model.Content{
		Reader: io.NopCloser(bytes.NewReader(b)),
		Size:   uint64(len(b)),
	})
	if err != nil {
		return fmt.Errorf("usecase set: %w", err)
	}

	return nil
}
