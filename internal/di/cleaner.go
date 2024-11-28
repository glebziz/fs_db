package di

import (
	"github.com/glebziz/fs_db/internal/usecase/cleaner"
)

func (c *Container) Cleaner() *cleaner.UseCase {
	if c.cleanerUseCase == nil {
		c.cleanerUseCase = cleaner.New(
			c.Core(),
			c.ContentRepo(),
			c.ContentFileRepo(),
			c.Badger(),
			c.DirRepo(),
			c.FileRepo(),
			c.Pool(),
			c.TransactionRepo(),
		)
	}

	return c.cleanerUseCase
}
