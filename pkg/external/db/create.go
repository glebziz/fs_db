package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamwriter"
)

func (db *db) Create(ctx context.Context, key string) (fs_db.File, error) {
	stream, err := db.client.SetFile(ctx)
	if err != nil {
		return nil, fmt.Errorf("set file: %w", grpc.ClientError(err))
	}

	err = stream.Send(&store.SetFileRequest{
		Data: &store.SetFileRequest_Header{
			Header: &store.FileHeader{
				Key: key,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("stream header send: %w", grpc.ClientError(err))
	}

	return streamwriter.New(store.ChunkSize_MAX, stream, func(p []byte) *store.SetFileRequest {
		return &store.SetFileRequest{
			Data: &store.SetFileRequest_Chunk{
				Chunk: p,
			},
		}
	}), nil
}