package db

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

func (db *db) SetReader(ctx context.Context, key string, reader io.Reader) error {
	err := db.container.Store().Set(ctx, key, reader)
	if err != nil {
		var errNotEnoughSpace model.NotEnoughSpaceError
		if errors.As(err, &errNotEnoughSpace) {
			errNotEnoughSpace.Close()
		}

		return fmt.Errorf("store usecase set: %w", err)
	}

	return nil
}
