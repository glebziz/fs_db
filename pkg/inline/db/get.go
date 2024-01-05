package db

import (
	"context"
	"fmt"
	"io"
)

func (db *db) Get(ctx context.Context, key string) ([]byte, error) {
	content, err := db.sUc.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("store usecase get: %w", err)
	}
	defer content.Reader.Close()

	b, err := io.ReadAll(content.Reader)
	if err != nil {
		return nil, fmt.Errorf("read all: %w", err)
	}

	return b, nil
}
