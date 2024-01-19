package file

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (r *rep) HardDelete(ctx context.Context, txId string, filter *model.FileFilter) (contentIds []string, err error) {
	query, args := r.prepareHardDeleteQuery(txId, filter)
	rows, err := r.p.DB(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var found bool
	for rows.Next() {
		found = true
		var (
			contentId sql.NullString
		)

		err = rows.Scan(&contentId)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		if contentId.Valid {
			contentIds = append(contentIds, contentId.String)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	if !found {
		return nil, fs_db.NotFoundErr
	}

	return contentIds, nil
}

func (_ *rep) prepareHardDeleteQuery(txId string, filter *model.FileFilter) (query string, args []interface{}) {
	args = []interface{}{txId}
	query = `delete from file
		where tx_id = $1`

	if filter != nil && filter.BeforeTs != nil {
		query = fmt.Sprintf(`%s and
				id in (
					select id
        			from (
            			select id, rank() over (partition by key order by ts desc) r
            			from file
						where tx_id = $1
                			and ts < $2
					)
					where r > 1
				)`, query)
		args = append(args, ptr.Val(filter.BeforeTs))
	}

	return fmt.Sprintf(`%s returning content_id`, query), args
}
