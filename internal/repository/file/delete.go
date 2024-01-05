package file

import (
	"context"
	"fmt"
	"time"
)

func (r *rep) Delete(ctx context.Context, txId, key string) error {
	res, err := r.p.DB(ctx).Exec(ctx, `
		insert into file(key, tx_id, ts)
		values ($1, $2, $3)`,
		key, txId, time.Now().UTC())
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
