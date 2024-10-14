package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Get(ctx context.Context, id string) (*model.ContentFile, error) {
	rows, err := r.p.DB(ctx).Query(ctx, `
		select id, parent_path
		from content_file
		where id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var file model.ContentFile

		err = rows.Scan(&file.Id, &file.Parent)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		return &file, err
	}

	return nil, fs_db.NotFoundErr
}
