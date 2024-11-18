package db

import (
	"context"
	"fmt"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (db *db) GetKeys(ctx context.Context) ([]string, error) {
	resp, err := db.client.GetKeys(ctx, &store.GetKeysRequest{})
	if err != nil {
		return nil, fmt.Errorf("get keys: %w", grpc.ClientError(err))
	}

	return resp.GetKeys(), nil
}
