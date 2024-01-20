package db

import (
	"context"
	"fmt"
	"io"

	store2 "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamreader"
)

type getWrapper struct {
	s store2.StoreV1_GetFileClient
}

func newGetWrapper(s store2.StoreV1_GetFileClient) *getWrapper {
	return &getWrapper{
		s: s,
	}
}

func (s *getWrapper) Recv() (streamreader.Request, error) {
	req, err := s.s.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv: %w", grpc.ClientError(err))
	}

	return req, nil
}

func (db *db) GetReader(ctx context.Context, key string) (io.ReadCloser, error) {
	stream, err := db.client.GetFile(ctx, &store2.GetFileRequest{
		Key: key,
	})
	if err != nil {
		return nil, fmt.Errorf("get file: %w", grpc.ClientError(err))
	}

	_, err = stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("recv header: %w", grpc.ClientError(err))
	}

	return io.NopCloser(streamreader.New(newGetWrapper(stream))), nil
}
