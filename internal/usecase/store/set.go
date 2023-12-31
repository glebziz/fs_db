package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Set(ctx context.Context, key string, content *model.Content) error {
	if key == "" {
		return fs_db.EmptyKeyErr
	}

	dir, err := u.dir.Select(ctx, content.Size)
	if err != nil {
		return fmt.Errorf("get parent dir: %w", err)
	}

	file := model.File{
		Id:         u.idGen.Generate(),
		Key:        key,
		ParentPath: dir.GetPath(),
	}

	err = u.fRepo.Store(ctx, file)
	if err != nil {
		return fmt.Errorf("file repository store: %w", err)
	}

	err = u.cRepo.Store(ctx, file.GetPath(), content)
	if err != nil {
		return fmt.Errorf("content repository store: %w", err)
	}

	return nil
}
