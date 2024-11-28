package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (db *db) GetKeys(ctx context.Context) ([]string, error) {
	resp, err := db.client.GetKeys(ctx, &store.GetKeysRequest{})
	if err != nil {
		return nil, fmt.Errorf("get keys: %w", errors.ClientError(err))
	}

	return resp.GetKeys(), nil
}
