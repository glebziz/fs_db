package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Get(ctx context.Context) ([]model.Dir, error) {
	rows, err := r.p.DB(ctx).Query(ctx, `
		select d.id, d.parent_path, coalesce(f.count, 0)
		from dir d
			left join (
			    select parent_path, count(*) count 
			    from content_file
			    group by parent_path
			) f on concat(d.parent_path, '/', d.id) = f.parent_path
		order by f.count`)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	dirs := make([]model.Dir, 0)

	for rows.Next() {
		var dir model.Dir

		err = rows.Scan(&dir.Id, &dir.ParentPath, &dir.FileCount)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		dirs = append(dirs, dir)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}

	return dirs, nil
}
