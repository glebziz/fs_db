package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) CommitTx(ctx context.Context, _ *store.CommitTxRequest) (*store.CommitTxResponse, error) {
	err := i.txUsecase.Commit(ctx)
	if err != nil {
		return nil, errors.Error(fmt.Errorf("tx usecase commit: %w", err))
	}

	return &store.CommitTxResponse{}, nil
}
