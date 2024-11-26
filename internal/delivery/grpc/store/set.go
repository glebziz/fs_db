package store

import (
	"fmt"

	"github.com/glebziz/fs_db"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamreader"
)

func (i *Service) SetFile(stream store.StoreV1_SetFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return grpc.Error(fmt.Errorf("stream recv: %w", err))
	}

	header := req.GetHeader()
	if header == nil {
		return grpc.Error(fs_db.HeaderNotFoundErr)
	}

	err = i.sUsecase.Set(stream.Context(), header.GetKey(), streamreader.New(stream))
	if err != nil {
		return grpc.Error(fmt.Errorf("store usecase set: %w", err))
	}

	err = stream.SendAndClose(&store.SetFileResponse{})
	if err != nil {
		return grpc.Error(fmt.Errorf("stream send and close: %w", err))
	}

	return nil
}
