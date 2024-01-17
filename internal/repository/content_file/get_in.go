package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) GetIn(ctx context.Context, ids []string) ([]model.ContentFile, error) {
	in, args := arrayArg(ids)
	rows, err := r.p.DB(ctx).Query(ctx, fmt.Sprintf(`
		select id, parent_path
		from content_file
		where id in (%s)`, in), args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var files []model.ContentFile
	for rows.Next() {
		var file model.ContentFile

		err = rows.Scan(&file.Id, &file.ParentPath)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		files = append(files, file)
	}

	return files, nil
}
