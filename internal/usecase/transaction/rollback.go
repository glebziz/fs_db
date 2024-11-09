package transaction

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Rollback(ctx context.Context) error {
	txId := model.GetTxId(ctx)
	_, err := u.txRepo.Delete(ctx, txId)
	if err != nil {
		return fmt.Errorf("tx repository delete: %w", err)
	}

	deleteFiles := u.fRepo.DeleteTx(ctx, txId)
	if len(deleteFiles) > 0 {
		u.cleaner.DeleteFilesAsync(ctx, deleteFiles)
	}
	return nil
}
