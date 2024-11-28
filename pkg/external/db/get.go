package db

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	errorsAdapter "github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (db *db) Get(ctx context.Context, key string) ([]byte, error) {
	stream, err := db.client.GetFile(ctx, &store.GetFileRequest{
		Key: key,
	})
	if err != nil {
		return nil, fmt.Errorf("get file: %w", errorsAdapter.ClientError(err))
	}

	_, err = stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv header: %w", errorsAdapter.ClientError(err))
	}

	var (
		buf  bytes.Buffer
		resp *store.GetFileResponse
	)
	for {
		resp, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("recv: %w", errorsAdapter.ClientError(err))
		}

		buf.Write(resp.GetChunk())
	}

	return buf.Bytes(), nil
}
