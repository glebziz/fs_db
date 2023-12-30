package content

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/glebziz/fs_db/internal/model"
	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func (r *rep) Get(_ context.Context, path string) (*model.Content, error) {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, pkgModel.NotFoundErr
	} else if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	st, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	return &model.Content{
		Reader: f,
		Size:   uint64(st.Size()),
	}, nil
}
