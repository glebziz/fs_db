package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func (r *rep) Get(ctx context.Context, key string) (*model.File, error) {
	rows, err := r.p.DB(ctx).Query(ctx, `
		select id, key, parent_path
		from file
		where key = $1`, key)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var file model.File

		err = rows.Scan(&file.Id, &file.Key, &file.ParentPath)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		return &file, err
	}

	return nil, pkgModel.NotFoundErr
}
