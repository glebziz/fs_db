package db

import (
	"context"
	"fmt"
)

func (db *db) Delete(ctx context.Context, key string) error {
	err := db.usecase.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("usecase delete: %w", err)
	}

	return nil
}
