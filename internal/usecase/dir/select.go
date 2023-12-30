package dir

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func (u *useCase) Select(ctx context.Context, size uint64) (*model.Dir, error) {
	rootMap, err := u.root.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("root get: %w", err)
	}

	dirs, err := u.dRepo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("dir repository get: %w", err)
	}

	for _, d := range dirs {
		root, ok := rootMap[d.ParentPath]
		if !ok {
			continue
		}

		if root.Free < size {
			delete(rootMap, d.ParentPath)
			continue
		}

		if d.FileCount < u.maxCount {
			return &d, nil
		}
	}

	root := rootMap.Select(size)
	if root == nil {
		return nil, pkgModel.SizeErr
	}

	d := model.Dir{
		Id:         u.idGen.Generate(),
		ParentPath: root.Path,
	}

	err = u.dRepo.Create(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &d, nil
}
