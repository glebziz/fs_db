package db

import (
	"context"
	"fmt"
)

func (db *db) GetKeys(ctx context.Context) ([]string, error) {
	keys, err := db.container.Store().GetKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("store usecase get keys: %w", err)
	}

	return keys, nil
}
