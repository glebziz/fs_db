package dir

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) GetRoots(_ context.Context) ([]model.Root, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	roots := make([]model.Root, 0, len(r.roots))
	for _, root := range r.roots {
		roots = append(roots, model.Root{
			Path:  root,
			Count: r.counts[root],
		})
	}

	return roots, nil
}
