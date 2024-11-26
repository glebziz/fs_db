package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/usecase/transaction"
)

type tx struct {
	id string

	txUc *transaction.UseCase
}

func (t *tx) ctx(ctx context.Context) context.Context {
	return model.StoreTxId(ctx, t.id)
}

func (t *tx) Commit(ctx context.Context) error {
	err := t.txUc.Commit(t.ctx(ctx))
	if err != nil {
		return fmt.Errorf("tx usecase commit: %w", err)
	}

	return nil
}

func (t *tx) Rollback(ctx context.Context) error {
	err := t.txUc.Rollback(t.ctx(ctx))
	if err != nil {
		return fmt.Errorf("tx usecase rollback: %w", err)
	}

	return nil
}
