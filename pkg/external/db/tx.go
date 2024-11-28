package db

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/glebziz/fs_db/internal/adapter/errors"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/grpc/interceptors/server"
)

type tx struct {
	id string

	client store.StoreV1Client
}

func (t *tx) ctx(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, server.TxIdKey, t.id)
}

func (t *tx) Commit(ctx context.Context) error {
	_, err := t.client.CommitTx(t.ctx(ctx), &store.CommitTxRequest{})
	if err != nil {
		return fmt.Errorf("commit tx: %w", errors.ClientError(err))
	}

	return nil
}

func (t *tx) Rollback(ctx context.Context) error {
	_, err := t.client.RollbackTx(t.ctx(ctx), &store.RollbackTxRequest{})
	if err != nil {
		return fmt.Errorf("rollback tx: %w", errors.ClientError(err))
	}

	return nil
}
