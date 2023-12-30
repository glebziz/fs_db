package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Get(ctx context.Context, key string) (*model.Content, error) {
	f, err := u.fRepo.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("file repository get: %w", err)
	}

	content, err := u.cRepo.Get(ctx, f.GetPath())
	if err != nil {
		return nil, fmt.Errorf("content repository get: %w", err)
	}

	return content, nil
}
