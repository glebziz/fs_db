package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
)

func (db *db) Open(ctx context.Context, key string) (fs_db.ReadFile, error) {
	content, err := db.container.Store().Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("store usecase get: %w", err)
	}

	return content, nil
}
