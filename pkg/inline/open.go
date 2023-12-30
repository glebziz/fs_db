package inline

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/pkg/inline/db"
	"github.com/glebziz/fs_db/pkg/model"
)

func Open(ctx context.Context, cfg *config.Storage) (model.DB, error) {
	b, err := db.New(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("new inline db: %w", err)
	}

	return b, nil
}
