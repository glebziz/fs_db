package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Rollback(ctx context.Context) error {
	txId := model.GetTxId(ctx)
	_, err := u.txRepo.Delete(ctx, txId)
	if err != nil {
		return fmt.Errorf("tx repository delete: %w", err)
	}

	contentIds, err := u.fRepo.HardDelete(ctx, txId, nil)
	if err != nil && !errors.Is(err, fs_db.NotFoundErr) {
		return fmt.Errorf("file repository delete by tx: %w", err)
	}

	u.cleaner.Send(contentIds)
	return nil
}
