package db

import (
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

func (db *db) SetReader(ctx context.Context, key string, reader io.Reader, size uint64) error {
	err := db.sUc.Set(ctx, key, &model.Content{
		Reader: io.NopCloser(reader),
		Size:   size,
	})
	if err != nil {
		return fmt.Errorf("store usecase set: %w", err)
	}

	return nil
}
