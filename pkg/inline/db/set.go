package db

import (
	"bytes"
	"context"
	"fmt"
)

func (db *db) Set(ctx context.Context, key string, b []byte) error {
	err := db.sUc.Set(ctx, key, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("store usecase set: %w", err)
	}

	return nil
}
