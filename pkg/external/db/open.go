package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/adapter/errors"
	"github.com/glebziz/fs_db/internal/adapter/seek_direction"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/dstreamreader"
)

func (db *db) Open(ctx context.Context, key string) (fs_db.ReadFile, error) {
	stream, err := db.client.GetFileV2(ctx)
	if err != nil {
		return nil, fmt.Errorf("get file v2: %w", errors.ClientError(err))
	}

	err = stream.Send(&store.GetFileV2Request{
		Data: &store.GetFileV2Request_Header{
			Header: &store.FileHeader{
				Key: key,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("send header: %w", errors.ClientError(err))
	}

	resp, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv size: %w", errors.ClientError(err))
	}

	return dstreamreader.New(stream, func(seek model.Seek) *store.GetFileV2Request {
		return &store.GetFileV2Request{
			Data: &store.GetFileV2Request_Pos{
				Pos: &store.Seek{
					Offset: seek.Pos,
					Dir:    seek_direction.ConvertToGrpc(seek.Dir),
				},
			},
		}
	}, resp.GetSize()), nil
}
