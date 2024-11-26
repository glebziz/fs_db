package di

import (
	"github.com/glebziz/fs_db/internal/delivery/grpc/store"
)

func (c *Container) StoreService() *store.Service {
	if c.storeService == nil {
		c.storeService = store.New(
			c.Store(),
			c.Transaction(),
		)
	}

	return c.storeService
}
