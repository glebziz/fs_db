package di

import (
	"github.com/glebziz/fs_db/internal/usecase/store"
)

func (c *Container) Store() *store.UseCase {
	if c.storeUseCase == nil {
		c.storeUseCase = store.New(
			c.Dir(),
			c.ContentRepo(),
			c.ContentFileRepo(),
			c.Core(),
			c.TransactionRepo(),
			c.Gen(),
			c.Rand(),
		)
	}

	return c.storeUseCase
}
