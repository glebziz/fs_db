package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) DeleteFile(ctx context.Context, req *store.DeleteFileRequest) (*store.DeleteFileResponse, error) {
	err := i.sUsecase.Delete(ctx, req.GetKey())
	if err != nil {
		return nil, errors.Error(fmt.Errorf("store usecase delete: %w", err))
	}

	return &store.DeleteFileResponse{}, nil
}
