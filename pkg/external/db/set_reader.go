package db

import (
	"context"
	"fmt"
	"io"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (db *db) SetReader(ctx context.Context, key string, reader io.Reader) error {
	stream, err := db.client.SetFile(ctx)
	if err != nil {
		return fmt.Errorf("set file: %w", grpc.ClientError(err))
	}

	err = stream.Send(&store.SetFileRequest{
		Data: &store.SetFileRequest_Header{
			Header: &store.FileHeader{
				Key: key,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("stream header send: %w", grpc.ClientError(err))
	}

	buf := make([]byte, store.ChunkSize_MAX)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}

		err = stream.Send(&store.SetFileRequest{
			Data: &store.SetFileRequest_Chunk{
				Chunk: buf[:n],
			},
		})
		if err != nil {
			return fmt.Errorf("stream chunk send: %w", grpc.ClientError(err))
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("stream close and recv: %w", grpc.ClientError(err))
	}

	return nil
}
