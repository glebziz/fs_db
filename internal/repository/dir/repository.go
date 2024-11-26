package dir

import (
	"errors"
	"fmt"
	"path"
	"sync"

	"github.com/google/uuid"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/os"
)

const (
	mkdirPerm = 0750
)

type Repo struct {
	roots  []string
	dirs   map[string]model.Dir
	counts map[string]uint64

	m sync.RWMutex
}

func New(rootDirs []string) (*Repo, error) {
	r := Repo{
		roots: rootDirs,

		dirs:   make(map[string]model.Dir, len(rootDirs)),
		counts: make(map[string]uint64, len(rootDirs)),
	}

	for i, root := range rootDirs {
		root = path.Join(root)
		r.roots[i] = root

		entries, err := os.ReadDir(root)
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(root, mkdirPerm)
			if err != nil {
				return nil, fmt.Errorf("mkdir all: %w", err)
			}
		}
		if err != nil {
			return nil, fmt.Errorf("read dir: %w, root: %s", err, root)
		}

		for _, entry := range entries {
			if !entry.IsDir() || uuid.Validate(entry.Name()) != nil {
				continue
			}

			dir := model.Dir{
				Name: entry.Name(),
				Root: root,
			}

			r.dirs[dir.Path()] = dir
			r.counts[root]++
		}
	}

	return &r, nil
}
