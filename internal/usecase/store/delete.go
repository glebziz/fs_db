package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Delete(ctx context.Context, key string) error {
	txId := model.GetTxId(ctx)
	err := u.fRepo.Delete(ctx, txId, key)
	if err != nil {
		return fmt.Errorf("file repository delete: %w", err)
	}

	return nil
}
