package db

import (
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamwriter"
)

func (db *db) SetReader(ctx context.Context, key string, reader io.Reader) error {
	stream, err := db.client.SetFile(ctx)
	if err != nil {
		return fmt.Errorf("set file: %w", errors.ClientError(err))
	}

	err = stream.Send(&store.SetFileRequest{
		Data: &store.SetFileRequest_Header{
			Header: &store.FileHeader{
				Key: key,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("stream header send: %w", errors.ClientError(err))
	}

	sw := streamwriter.New(store.ChunkSize_MAX, stream, func(p []byte) *store.SetFileRequest {
		return &store.SetFileRequest{
			Data: &store.SetFileRequest_Chunk{
				Chunk: p,
			},
		}
	})
	defer sw.Close()

	_, err = io.Copy(sw, reader)
	if err != nil {
		return fmt.Errorf("copy: %w", errors.ClientError(err))
	}

	return nil
}
