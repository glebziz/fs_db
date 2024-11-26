package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/os"
)

func (r *Repo) Create(_ context.Context, dir model.Dir) error {
	err := os.MkdirAll(dir.Path(), mkdirPerm)
	if err != nil {
		return fmt.Errorf("mkdir all: %w", err)
	}

	r.m.Lock()
	defer r.m.Unlock()

	r.dirs[dir.Path()] = dir
	r.counts[dir.Root]++

	return nil
}
