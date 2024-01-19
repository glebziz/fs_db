package store

import (
	"fmt"
	"io"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamreader"
)

type wrapper struct {
	s store.StoreV1_SetFileServer
}

func newWrapper(s store.StoreV1_SetFileServer) *wrapper {
	return &wrapper{
		s: s,
	}
}

func (s *wrapper) Recv() (streamreader.Request, error) {
	return s.s.Recv()
}

func (i *implementation) SetFile(stream store.StoreV1_SetFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return grpc.Error(fmt.Errorf("stream recv: %w", err))
	}

	header := req.GetHeader()
	if header == nil {
		return grpc.Error(fs_db.HeaderNotFoundErr)
	}

	err = i.sUsecase.Set(stream.Context(), header.Key, &model.Content{
		Reader: io.NopCloser(streamreader.New(newWrapper(stream))),
		Size:   header.Size,
	})
	if err != nil {
		return grpc.Error(fmt.Errorf("store usecase set: %w", err))
	}

	err = stream.SendAndClose(&store.SetFileResponse{})
	if err != nil {
		return grpc.Error(fmt.Errorf("stream send and close: %w", err))
	}

	return nil
}
