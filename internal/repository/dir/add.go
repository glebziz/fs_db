package dir

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Add(_ context.Context, dir model.Dir) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.dirs[dir.Path()]
	if ok {
		return nil
	}

	r.counts[dir.Root]++
	r.dirs[dir.Path()] = dir

	return nil
}
