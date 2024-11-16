package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Get(ctx context.Context, id string) (model.ContentFile, error) {
	parent, err := r.p.DB(ctx).Get(r.key(id))
	if err != nil {
		return model.ContentFile{}, fmt.Errorf("db get: %w", err)
	}

	return model.ContentFile{
		Id:     id,
		Parent: string(parent),
	}, nil
}
