package dir

import (
	"context"
	"fmt"

	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func (r *rep) Delete(ctx context.Context, key string) error {
	res, err := r.p.DB(ctx).Exec(ctx, `
		delete from file
		where key = $1`, key)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if affected == 0 {
		return pkgModel.NotFoundErr
	}

	return nil
}
