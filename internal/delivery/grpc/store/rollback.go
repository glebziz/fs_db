package store

import (
	"context"
	"fmt"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (i *implementation) RollbackTx(ctx context.Context, _ *store.RollbackTxRequest) (*store.RollbackTxResponse, error) {
	err := i.txUsecase.Rollback(ctx)
	if err != nil {
		return nil, grpc.Error(fmt.Errorf("tx usecase rollback: %w", err))
	}

	return &store.RollbackTxResponse{}, nil
}
