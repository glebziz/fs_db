package store

import (
	"context"
	"fmt"

	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (i *Service) CommitTx(ctx context.Context, _ *store.CommitTxRequest) (*store.CommitTxResponse, error) {
	err := i.txUsecase.Commit(ctx)
	if err != nil {
		return nil, grpc.Error(fmt.Errorf("tx usecase commit: %w", err))
	}

	return &store.CommitTxResponse{}, nil
}
