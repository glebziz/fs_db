package file

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Delete(ctx context.Context, f model.File) error {
	err := r.p.DB(ctx).Delete(r.key(f.ContentId))
	if err != nil {
		return fmt.Errorf("db delete: %w", err)
	}

	return nil
}
