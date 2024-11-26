package di

import (
	"github.com/samber/lo"

	"github.com/glebziz/fs_db/internal/db/badger"
)

func (c *Container) Badger() *badger.Manager {
	if c.badger == nil {
		c.badger = lo.Must(badger.New(c.cfg.Storage.DbPath))
	}

	return c.badger
}
