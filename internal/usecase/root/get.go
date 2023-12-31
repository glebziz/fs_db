package root

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Get(ctx context.Context) (model.RootMap, error) {
	rootMap := make(model.RootMap, len(u.rootDirs))
	for _, path := range u.rootDirs {
		st, err := u.manager.Usage(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("usage: %w", err)
		}

		count, err := u.repo.CountByParent(ctx, path)
		if err != nil {
			return nil, fmt.Errorf("count by parent: %w", err)
		}

		rootMap[path] = &model.Root{
			Path:  path,
			Free:  st.Free,
			Count: count,
		}
	}

	return rootMap, nil
}
