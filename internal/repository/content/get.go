package content

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/utils/os"
)

func (r *Repo) Get(_ context.Context, path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fs_db.NotFoundErr
	} else if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	return f, nil
}
