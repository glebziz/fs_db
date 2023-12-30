package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/utils/grpc"

	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *implementation) DeleteFile(ctx context.Context, req *store.DeleteFileRequest) (*store.DeleteFileResponse, error) {
	err := i.usecase.Delete(ctx, req.Key)
	if err != nil {
		return nil, grpc.Error(fmt.Errorf("usecase delete: %w", err))
	}

	return &store.DeleteFileResponse{}, nil
}
