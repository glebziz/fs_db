package db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model/content"
)

func (db *db) Create(ctx context.Context, key string) (fs_db.File, error) {
	rw, contents := content.New()
	rw.Go(func() {
		err := db.container.Store().Set(ctx, key, contents)
		if err != nil {
			rw.SetError(fmt.Errorf("store usecase set: %w", err))
		}
	})

	return rw, nil
}
