package store

import (
	"context"
	"fmt"
)

func (u *useCase) Delete(ctx context.Context, key string) error {
	file, err := u.fRepo.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("file repository get: %w", err)
	}

	err = u.cRepo.Delete(ctx, file.GetPath())
	if err != nil {
		return fmt.Errorf("content repository delete: %w", err)
	}

	err = u.fRepo.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("file repository delete: %w", err)
	}

	return nil
}
