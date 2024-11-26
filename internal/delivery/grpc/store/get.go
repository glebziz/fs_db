package store

import (
	"errors"
	"fmt"
	"io"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (i *Service) GetFile(req *store.GetFileRequest, stream store.StoreV1_GetFileServer) error {
	content, err := i.sUsecase.Get(stream.Context(), req.GetKey())
	if err != nil {
		return grpc.Error(fmt.Errorf("store usecase get: %w", err))
	}
	defer content.Close()

	err = stream.Send(&store.GetFileResponse{
		Data: &store.GetFileResponse_Header{
			Header: &store.FileHeader{
				Key: req.GetKey(),
			},
		},
	})
	if err != nil {
		return grpc.Error(fmt.Errorf("stream file header send: %w", err))
	}

	chunk := make([]byte, store.ChunkSize_MAX)
	for {
		var n int
		n, err = content.Read(chunk)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return grpc.Error(fmt.Errorf("read: %w", err))
		}

		err = stream.Send(&store.GetFileResponse{
			Data: &store.GetFileResponse_Chunk{
				Chunk: chunk[:n],
			},
		})
		if err != nil {
			return grpc.Error(fmt.Errorf("stream chunk send: %w", err))
		}
	}

	return nil
}
