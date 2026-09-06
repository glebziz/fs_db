package di

import (
	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/utils/must"
)

func (c *Container) Badger() *badger.Manager {
	if c.badger == nil {
		c.badger = must.Must(badger.New(c.cfg.Storage.DbPath))
	}

	return c.badger
}
