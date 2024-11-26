package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/async"
)

func (db *db) Create(ctx context.Context, key string) (fs_db.File, error) {
	rw := async.NewReadWriter()
	rw.Add(1)
	go func() {
		defer rw.Done()

		err := db.container.Store().Set(ctx, key, rw)
		if err != nil {
			var errNotEnoughSpace model.NotEnoughSpaceError
			if errors.As(err, &errNotEnoughSpace) {
				errNotEnoughSpace.Close()
			}

			rw.SetError(fmt.Errorf("store usecase set: %w", err))
		}
	}()

	return rw, nil
}
