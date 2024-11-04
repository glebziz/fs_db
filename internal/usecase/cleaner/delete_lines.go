package cleaner

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func (c *Cleaner) deleteLines(ctx context.Context) error {
	tx, err := c.txRepo.Oldest(ctx)
	if errors.Is(err, fs_db.TxNotFoundErr) {
		tx = &model.Transaction{
			Seq: sequence.Next(),
		}
	} else if err != nil {
		return fmt.Errorf("tx repo oldest: %w", err)
	}

	c.fRepo.DeleteOld(ctx, model.MainTxId, tx.Seq)
	if err != nil {
		return fmt.Errorf("file repo hard delete: %w", err)
	}

	//err = c.deleteContent(ctx, cIds)
	//if err != nil {
	//	return fmt.Errorf("delete lines: %w", err)
	//}

	return nil
}
