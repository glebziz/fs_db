package file

import (
	"context"
	"fmt"
	"time"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (r *rep) UpdateTx(ctx context.Context, oldTxId string, newTxId string, filter *model.FileFilter) error {
	query, args := r.prepareUpdateTxQuery(oldTxId, newTxId, filter)
	_, err := r.p.DB(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func (_ *rep) prepareUpdateTxQuery(oldTxId, newTxId string, filter *model.FileFilter) (query string, args []interface{}) {
	if filter == nil || filter.BeforeTs == nil {
		return `
			update file
			set tx_id = $1, ts = $2
			where tx_id = $3
		`, []interface{}{newTxId, time.Now().UTC(), oldTxId}
	}

	return `
		with err_keys as (
			select m.key
			from file m
				inner join file c on m.key = c.key and c.tx_id = $1
			where m.tx_id = $2
			group by m.key
			having max(m.ts) > $3
		)
		update file
		set tx_id = $2, ts = $4
		from (
			select count(err_keys.key) cnt
			from err_keys
		) err_count
		where tx_id = $1 and cnt = 0
	`, []interface{}{oldTxId, newTxId, ptr.Val(filter.BeforeTs), time.Now().UTC()}
}
