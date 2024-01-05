package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Store(ctx context.Context, file model.ContentFile) error {
	res, err := r.p.DB(ctx).Exec(ctx, `
		insert into content_file(id, parent_path)
		values ($1, $2)`,
		file.Id, file.ParentPath)
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
