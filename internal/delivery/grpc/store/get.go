package store

import (
	"fmt"
	"io"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (i *implementation) GetFile(req *store.GetFileRequest, stream store.StoreV1_GetFileServer) error {
	content, err := i.usecase.Get(stream.Context(), req.Key)
	if err != nil {
		return grpc.Error(fmt.Errorf("usecase get: %w", err))
	}
	defer content.Reader.Close()

	chunk := make([]byte, store.ChunkSize_MAX)
	for {
		n, err := content.Reader.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return grpc.Error(fmt.Errorf("read: %w", err))
		}

		err = stream.Send(&store.GetFileResponse{
			Chunk: chunk[:n],
		})
		if err != nil {
			return grpc.Error(fmt.Errorf("stream chunk send: %w", err))
		}
	}

	return nil
}
