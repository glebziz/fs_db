package dir

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/google/uuid"
	"github.com/shirou/gopsutil/disk"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source repository.go -package mocks -destination mocks/mocks.go -typed true

type osAdapter interface {
	Usage(ctx context.Context, path string) (*disk.UsageStat, error)
	MkdirAll(ctx context.Context, path string, perm os.FileMode) error
	ReadDir(ctx context.Context, name string) ([]os.DirEntry, error)
}

const (
	mkdirPerm os.FileMode = 0750
)

type Repo struct {
	roots  []string
	dirs   map[string]model.Dir
	counts map[string]uint64

	m sync.RWMutex

	os osAdapter
}

func New(ctx context.Context, rootDirs []string, osa osAdapter) (*Repo, error) {
	r := Repo{
		roots: rootDirs,

		dirs:   make(map[string]model.Dir, len(rootDirs)),
		counts: make(map[string]uint64, len(rootDirs)),

		os: osa,
	}

	for i, root := range rootDirs {
		root = path.Join(root)
		r.roots[i] = root

		entries, err := osa.ReadDir(ctx, root)
		if errors.Is(err, os.ErrNotExist) {
			err = osa.MkdirAll(ctx, root, mkdirPerm)
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
