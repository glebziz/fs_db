package transaction

import (
	"context"
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

	var filter model.FileFilter
	switch tx.IsoLevel {
	case fs_db.IsoLevelReadUncommitted,
		fs_db.IsoLevelReadCommitted:
	case fs_db.IsoLevelRepeatableRead,
		fs_db.IsoLevelSerializable:
		filter.BeforeSeq = ptr.Ptr(tx.Seq)
	}

	err = u.fRepo.UpdateTx(ctx, txId, model.MainTxId, filter)
	if err != nil {
		return fmt.Errorf("file repository update tx: %w", err)
	}

	return nil
}
