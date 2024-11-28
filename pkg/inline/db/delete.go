package db

import (
	"context"
	"fmt"
)

func (db *db) Delete(ctx context.Context, key string) error {
	err := db.container.Store().Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("store usecase delete: %w", err)
	}

	return nil
}
