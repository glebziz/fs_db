package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (u *useCase) Commit(ctx context.Context) error {
	txId := model.GetTxId(ctx)
	tx, err := u.txRepo.Delete(ctx, txId)
	if err != nil {
		return fmt.Errorf("tx repository delete: %w", err)
	}

	var filter *model.FileFilter
	switch tx.IsoLevel {
	case fs_db.IsoLevelReadUncommitted,
		fs_db.IsoLevelReadCommitted:
	case fs_db.IsoLevelRepeatableRead,
		fs_db.IsoLevelSerializable:
		filter = &model.FileFilter{
			BeforeTs: ptr.Ptr(tx.CreateTs),
		}
	}

	err = u.fRepo.UpdateTx(ctx, txId, model.MainTxId, filter)
	if err != nil {
		return fmt.Errorf("file repository update tx: %w", err)
	}

	if filter != nil {
		contentIds, err := u.fRepo.HardDelete(ctx, txId, nil)
		if errors.Is(err, fs_db.NotFoundErr) {
			return nil
		} else if err != nil {
			return fmt.Errorf("file repository delete by tx: %w", err)
		}

		u.cleaner.Send(contentIds)
		return fs_db.TxSerializationErr
	}

	return nil
}
