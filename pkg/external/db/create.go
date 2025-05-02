package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/adapter/errors"
	"github.com/glebziz/fs_db/internal/adapter/seek_direction"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamwriter"
)

func (db *db) Create(ctx context.Context, key string) (fs_db.File, error) {
	stream, err := db.client.SetFileV2(ctx)
	if err != nil {
		return nil, fmt.Errorf("set file: %w", errors.ClientError(err))
	}

	err = stream.Send(&store.SetFileV2Request{
		Data: &store.SetFileV2Request_Header{
			Header: &store.FileHeader{
				Key: key,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("stream header send: %w", errors.ClientError(err))
	}

	return streamwriter.New(store.ChunkSize_MAX, stream, func(p []byte, seek model.Seek) *store.SetFileV2Request {
		return &store.SetFileV2Request{
			Data: &store.SetFileV2Request_Chunk{
				Chunk: &store.DataChunk{
					Chunk: p,
					Pos: &store.Seek{
						Offset: seek.Pos,
						Dir:    seek_direction.ConvertToGrpc(seek.Dir),
					},
				},
			},
		}
	}), nil
}
