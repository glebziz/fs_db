package file

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func (r *rep) Get(ctx context.Context, txId, key string, filter *model.FileFilter) (*model.File, error) {
	query, args := r.prepareGetQuery(txId, key, filter)
	rows, err := r.p.DB(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var contentId sql.NullString

		err = rows.Scan(&key, &contentId)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		if contentId.Valid {
			return &model.File{
				Key:       key,
				ContentId: contentId.String,
			}, nil
		}
	}

	return nil, fs_db.NotFoundErr
}

func (_ *rep) prepareGetQuery(txId, key string, filter *model.FileFilter) (query string, args []interface{}) {
	args = []interface{}{key, txId}
	mainQuery := `select key, content_id, ts
		from file
		where key = $1 and
		      tx_id = $2`

	unionQuery := `select key, content_id, ts
		from file
		where key = $1`

	if filter != nil {
		if filter.TxId != nil {
			unionQuery = fmt.Sprintf(`%s and
				tx_id = $3`, unionQuery)
			args = append(args, ptr.Val(filter.TxId))
		}

		if filter.BeforeTs != nil {
			unionQuery = fmt.Sprintf(`%s and
				ts < $4`, unionQuery)
			args = append(args, ptr.Val(filter.BeforeTs))
		}
	}

	return fmt.Sprintf(`
		select key, content_id
		from (
			%s
			union
			%s
			order by ts desc
			limit 1
		) res`,
		mainQuery, unionQuery), args
}
