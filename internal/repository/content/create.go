package content

import (
	"context"
	"fmt"
	"os"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Create(_ context.Context, path string) (model.ReadWriteSeekCloser, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return f, err
}
