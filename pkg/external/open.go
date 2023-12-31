package external

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/pkg/external/db"
)

func Open(ctx context.Context, url string) (fs_db.DB, error) {
	b, err := db.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("new db: %w", err)
	}

	return b, nil
}
