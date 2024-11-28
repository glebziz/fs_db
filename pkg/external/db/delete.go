package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (db *db) Delete(ctx context.Context, key string) error {
	_, err := db.client.DeleteFile(ctx, &store.DeleteFileRequest{
		Key: key,
	})
	if err != nil {
		return fmt.Errorf("delete file: %w", errors.ClientError(err))
	}

	return nil
}
