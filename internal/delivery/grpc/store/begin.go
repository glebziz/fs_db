package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	isoLevel "github.com/glebziz/fs_db/internal/adapter/iso_level"
	store "github.com/glebziz/fs_db/internal/proto"
)

func (i *Service) BeginTx(ctx context.Context, req *store.BeginTxRequest) (*store.BeginTxResponse, error) {
	txId, err := i.txUsecase.Begin(ctx, isoLevel.Convert(req.GetIsoLevel()))
	if err != nil {
		return nil, errors.Error(fmt.Errorf("tx usecase begin: %w", err))
	}

	return &store.BeginTxResponse{
		Id: txId,
	}, nil
}
