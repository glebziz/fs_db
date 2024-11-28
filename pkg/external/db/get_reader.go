package db

import (
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamreader"
)

func (db *db) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	stream, err := db.client.GetFile(ctx, &store.GetFileRequest{
		Key: key,
	})
	if err != nil {
		return nil, fmt.Errorf("get file: %w", errors.ClientError(err))
	}

	_, err = stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv header: %w", errors.ClientError(err))
	}

	return io.NopCloser(streamreader.New(stream)), nil
}
