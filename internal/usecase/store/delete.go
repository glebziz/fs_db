package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Delete(ctx context.Context, key string) error {
	txId := model.GetTxId(ctx)

	err := u.fRepo.Store(ctx, model.File{
		Key:       key,
		TxId:      txId,
		ContentId: u.idGen.Generate(),
	})
	if err != nil {
		return fmt.Errorf("file repository store: %w", err)
	}

	return nil
}
