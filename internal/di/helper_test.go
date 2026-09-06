package di

import (
	"context"
	"path"
	"testing"

	"github.com/glebziz/fs_db/config"
)

type prepareFunc func(c *Container)

func newContainer(t *testing.T) *Container {
	dir := t.TempDir()
	c := New(context.Background(), config.Config{
		Storage: config.Storage{
			DbPath:      path.Join(dir, "db"),
			MaxDirCount: 10,
			RootDirs:    []string{path.Join(dir, "root")},
		},
	})

	t.Cleanup(func() {
		c.Pool().Stop()
		c.Badger().Close()
	})

	return c
}
