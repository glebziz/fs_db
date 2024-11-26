package di

import (
	"github.com/glebziz/fs_db/internal/usecase/core"
)

func (c *Container) Core() *core.UseCase {
	if c.coreUseCase == nil {
		c.coreUseCase = core.New(c.FileRepo())
	}

	return c.coreUseCase
}
