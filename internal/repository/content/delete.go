package content

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/glebziz/fs_db"
)

func (r *rep) Delete(_ context.Context, path string) error {
	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return fs_db.NotFoundErr
	} else if err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	return nil
}
