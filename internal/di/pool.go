package di

import (
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

func (c *Container) Pool() *wpool.Pool {
	if c.pool == nil {
		c.pool = wpool.New(wpool.Options{
			NumWorkers:   c.cfg.WPool.NumWorkers,
			SendDuration: c.cfg.WPool.SendDuration,
		})
	}

	return c.pool
}
