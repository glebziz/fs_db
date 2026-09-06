package content

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Create(ctx context.Context, path string) (model.ReadWriteSeekCloser, error) {
	f, err := r.os.Create(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return f, err
}
