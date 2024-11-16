package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/utils/grpc"

	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *implementation) DeleteFile(ctx context.Context, req *store.DeleteFileRequest) (*store.DeleteFileResponse, error) {
	err := i.sUsecase.Delete(ctx, req.GetKey())
	if err != nil {
		return nil, grpc.Error(fmt.Errorf("store usecase delete: %w", err))
	}

	return &store.DeleteFileResponse{}, nil
}
