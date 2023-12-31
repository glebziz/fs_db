package content

import (
	"context"
	"errors"
	"fmt"
	"os"

	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func (r *rep) Delete(_ context.Context, path string) error {
	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return pkgModel.NotFoundErr
	} else if err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	return nil
}
