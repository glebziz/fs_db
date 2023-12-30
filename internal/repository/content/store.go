package content

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Store(_ context.Context, path string, content *model.Content) error {
	err := os.MkdirAll(filepath.Dir(path), 0666)
	if err != nil {
		return fmt.Errorf("mkdir all: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, content.Reader)
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}
