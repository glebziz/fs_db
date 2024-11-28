package file

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) GetAll(ctx context.Context) ([]model.File, error) {
	items, err := r.p.DB(ctx).GetAll(r.key(""))
	if err != nil {
		return nil, err
	}

	files := make([]model.File, len(items))
	for i, item := range items {
		err = unmarshalFile(item.Value, &files[i])
		if err != nil {
			return nil, err
		}
	}

	return files, nil
}
