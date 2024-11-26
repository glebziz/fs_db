package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/disk"
	"github.com/glebziz/fs_db/internal/utils/os"
)

func (r *Repo) Get(ctx context.Context) (model.Dirs, error) {
	r.m.RLock()

	dirs := make(model.Dirs, 0, len(r.dirs))
	for _, dir := range r.dirs {
		dirs = append(dirs, dir)
	}

	r.m.RUnlock()

	freeByRoots := make(map[string]uint64, len(r.roots))
	for _, root := range r.roots {
		stat, err := disk.Usage(ctx, root)
		if err != nil {
			return nil, fmt.Errorf("disk usage: %w, root: %s", err, root)
		}

		freeByRoots[root] = stat.Free
	}

	for i := range dirs {
		path := dirs[i].Path()

		entities, err := os.ReadDir(path)
		if err != nil {
			return nil, fmt.Errorf("read dir: %w, dir: %s", err, path)
		}

		dirs[i].Free = freeByRoots[dirs[i].Root]
		dirs[i].Count = uint64(len(entities))
	}

	return dirs, nil
}
