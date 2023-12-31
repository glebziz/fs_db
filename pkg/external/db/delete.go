package db

import (
	"context"
	"fmt"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (db *db) Delete(ctx context.Context, key string) error {
	_, err := db.client.DeleteFile(ctx, &store.DeleteFileRequest{
		Key: key,
	})
	if err != nil {
		return fmt.Errorf("delete file: %w", grpc.ClientError(err))
	}

	return nil
}
