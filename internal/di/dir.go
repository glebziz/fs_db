package di

import (
	"github.com/glebziz/fs_db/internal/usecase/dir"
)

func (c *Container) Dir() *dir.UseCase {
	if c.dirUseCase == nil {
		c.dirUseCase = dir.New(
			c.cfg.Storage.MaxDirCount,
			c.DirRepo(),
			c.Gen(),
		)
	}

	return c.dirUseCase
}
