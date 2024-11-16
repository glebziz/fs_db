package store

import (
	"context"
	"fmt"

	isoLevel "github.com/glebziz/fs_db/internal/adapter/iso_level"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc"
)

func (i *implementation) BeginTx(ctx context.Context, req *store.BeginTxRequest) (*store.BeginTxResponse, error) {
	txId, err := i.txUsecase.Begin(ctx, isoLevel.Convert(req.GetIsoLevel()))
	if err != nil {
		return nil, grpc.Error(fmt.Errorf("tx usecase begin: %w", err))
	}

	return &store.BeginTxResponse{
		Id: txId,
	}, nil
}
