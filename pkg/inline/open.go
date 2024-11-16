package inline

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/pkg/inline/db"
)

// Open returns the client for the inlined fs db.
func Open(ctx context.Context, cfg config.Config) (fs_db.DB, error) {
	b, err := db.New(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("new inline db: %w", err)
	}

	return b, nil
}
