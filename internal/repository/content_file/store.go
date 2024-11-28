package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Store(ctx context.Context, file model.ContentFile) error {
	err := r.p.DB(ctx).Set(r.key(file.Id), []byte(file.Parent))
	if err != nil {
		return fmt.Errorf("db set: %w", err)
	}

	return nil
}
