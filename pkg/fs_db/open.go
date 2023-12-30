package fs_db

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/pkg/fs_db/db"
	"github.com/glebziz/fs_db/pkg/model"
)

func Open(ctx context.Context, url string) (model.DB, error) {
	b, err := db.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("new db: %w", err)
	}

	return b, nil
}
