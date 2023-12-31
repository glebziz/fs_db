package db

import (
	"bytes"
	"context"
	"fmt"
)

func (db *db) Set(ctx context.Context, key string, b []byte) error {
	err := db.SetReader(ctx, key, bytes.NewReader(b), uint64(len(b)))
	if err != nil {
		return fmt.Errorf("set reader: %w", err)
	}

	return nil
}
