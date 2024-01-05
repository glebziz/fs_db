package file

import (
	"context"
	"fmt"
	"time"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Store(ctx context.Context, txId string, file model.File) error {
	res, err := r.p.DB(ctx).Exec(ctx, `
		insert into file(key, content_id, tx_id, ts)
		values ($1, $2, $3, $4)`,
		file.Key, file.ContentId, txId, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("no rows are affected")
	}

	return nil
}
