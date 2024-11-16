package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Get(ctx context.Context) (model.Dirs, error) {
	roots, err := u.dRepo.GetRoots(ctx)
	if err != nil {
		return nil, fmt.Errorf("get roots: %w", err)
	}

	for _, root := range roots {
		if root.Count > 0 {
			continue
		}

		err = u.dRepo.Create(ctx, model.Dir{
			Name: u.nameGen.Generate(),
			Root: root.Path,
		})
		if err != nil {
			return nil, fmt.Errorf("create: %w", err)
		}
	}

	dirs, err := u.dRepo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	for i, dir := range dirs {
		if dir.Count < u.maxCount {
			continue
		}

		err = u.dRepo.Remove(ctx, dir)
		if err != nil {
			return nil, fmt.Errorf("remove: %w", err)
		}

		dirs[i].Count = 0
		dirs[i].Name = u.nameGen.Generate()
		err = u.dRepo.Create(ctx, dirs[i])
		if err != nil {
			return nil, fmt.Errorf("create: %w", err)
		}
	}

	return dirs, nil
}
