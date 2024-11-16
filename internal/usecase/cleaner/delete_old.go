package cleaner

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (u *useCase) DeleteOld(ctx context.Context) error {
	tx, err := u.txRepo.Oldest(ctx)
	if errors.Is(err, fs_db.TxNotFoundErr) {
		tx = model.Transaction{
			Seq: sequence.Next(),
		}
	} else if err != nil {
		return fmt.Errorf("tx repo oldest: %w", err)
	}

	files := u.core.DeleteOld(ctx, model.MainTxId, tx.Seq)
	err = u.DeleteFiles(ctx, files)
	if err != nil {
		return fmt.Errorf("delete files: %w", err)
	}

	return nil
}
