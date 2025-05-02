package db

import (
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

func (db *db) SetReader(ctx context.Context, key string, reader io.Reader) error {
	err := db.container.Store().Set(ctx, key, model.SingleContent(reader))
	if err != nil {
		return fmt.Errorf("store usecase set: %w", err)
	}

	return nil
}
