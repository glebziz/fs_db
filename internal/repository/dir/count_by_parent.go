package dir

import (
	"context"
	"fmt"
)

func (r *rep) CountByParent(ctx context.Context, parent string) (uint64, error) {
	rows, err := r.p.DB(ctx).Query(ctx, `
		select count(*)
		from dir
		where parent_path = $1`, parent)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var count uint64

		err = rows.Scan(&count)
		if err != nil {
			return 0, fmt.Errorf("scan: %w", err)
		}

		return count, nil
	}

	err = rows.Err()
	if err != nil {
		return 0, fmt.Errorf("rows: %w", err)
	}

	return 0, nil
}
