package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) RollbackTx(ctx context.Context, _ *store.RollbackTxRequest) (*store.RollbackTxResponse, error) {
	err := i.txUsecase.Rollback(ctx)
	if err != nil {
		return nil, errors.Error(fmt.Errorf("tx usecase rollback: %w", err))
	}

	return &store.RollbackTxResponse{}, nil
}
