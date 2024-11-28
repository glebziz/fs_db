package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) GetKeys(ctx context.Context, _ *store.GetKeysRequest) (*store.GetKeysResponse, error) {
	keys, err := i.sUsecase.GetKeys(ctx)
	if err != nil {
		return nil, errors.Error(fmt.Errorf("store usecase get keys: %w", err))
	}

	return &store.GetKeysResponse{
		Keys: keys,
	}, nil
}
