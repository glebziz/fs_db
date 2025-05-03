package content

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/glebziz/fs_db"
)

func (r *Repo) Delete(ctx context.Context, path string) error {
	err := r.os.Remove(ctx, path)
	if errors.Is(err, os.ErrNotExist) {
		return fs_db.ErrNotFound
	} else if err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	return nil
}
