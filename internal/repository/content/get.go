package content

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Get(ctx context.Context, path string) (model.ReadSeekCloser, error) {
	f, err := r.os.Open(ctx, path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fs_db.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	return f, nil
}
