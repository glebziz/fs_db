package cleaner

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (c *Cleaner) deleteLines(ctx context.Context) error {
	tx, err := c.txRepo.Oldest(ctx)
	if errors.Is(err, fs_db.TxNotFoundErr) {
		tx = &model.Transaction{
			CreateTs: time.Now(),
		}
	} else if err != nil {
		return fmt.Errorf("tx repo oldest: %w", err)
	}

	cIds, err := c.fRepo.HardDelete(ctx, model.MainTxId, &model.FileFilter{
		BeforeTs: ptr.Ptr(tx.CreateTs),
	})
	if err != nil {
		return fmt.Errorf("file repo hard delete: %w", err)
	}

	err = c.deleteContent(ctx, cIds)
	if err != nil {
		return fmt.Errorf("delete lines: %w", err)
	}

	return nil
}
