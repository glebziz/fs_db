package cleaner

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/glebziz/fs_db"
)

func (c *Cleaner) deleteContent(ctx context.Context, contentIds []string) error {
	if len(contentIds) == 0 {
		return nil
	}

	cfs, err := c.cfRepo.GetIn(ctx, contentIds)
	if err != nil {
		return fmt.Errorf("content file repo get in: %w", err)
	}

	if cfs == nil {
		return nil
	}

	for _, cf := range cfs {
		err = c.cRepo.Delete(ctx, cf.GetPath())
		if errors.Is(err, fs_db.NotFoundErr) {
			slog.Warn("content not exists",
				slog.String("path", cf.GetPath()),
			)
		} else if err != nil {
			return fmt.Errorf("content repo delete: %w", err)
		}
	}

	err = c.cfRepo.Delete(ctx, contentIds)
	if err != nil {
		return fmt.Errorf("content file repo delete: %w", err)
	}

	return nil
}
