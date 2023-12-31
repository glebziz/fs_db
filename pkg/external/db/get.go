package db

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (db *db) Get(ctx context.Context, key string) ([]byte, error) {
	stream, err := db.client.GetFile(ctx, &store.GetFileRequest{
		Key: key,
	})
	if err != nil {
		return nil, fmt.Errorf("get file: %w", grpc.ClientError(err))
	}

	var buf bytes.Buffer
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("recv: %w", grpc.ClientError(err))
		}

		buf.Write(req.GetChunk())
	}

	return buf.Bytes(), nil
}
