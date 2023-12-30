package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Store(ctx context.Context, file model.File) error {
	res, err := r.p.DB(ctx).Exec(ctx, `
		insert into file(id, key, parent_path)
		values ($1, $2, $3)
		on conflict (key) do update
		set id = excluded.id, parent_path = excluded.parent_path`,
		file.Id, file.Key, file.ParentPath)
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
