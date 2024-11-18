package store

import (
	"context"
	"fmt"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (i *implementation) GetKeys(ctx context.Context, _ *store.GetKeysRequest) (*store.GetKeysResponse, error) {
	keys, err := i.sUsecase.GetKeys(ctx)
	if err != nil {
		return nil, grpc.Error(fmt.Errorf("store usecase get keys: %w", err))
	}

	return &store.GetKeysResponse{
		Keys: keys,
	}, nil
}
