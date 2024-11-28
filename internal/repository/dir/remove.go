package dir

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *Repo) Remove(_ context.Context, dir model.Dir) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.dirs[dir.Path()]
	if !ok {
		return nil
	}

	r.counts[dir.Root]--
	delete(r.dirs, dir.Path())

	return nil
}
