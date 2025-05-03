package di

import (
	"github.com/samber/lo"

	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
)

func (c *Container) DirRepo() *dirRepo.Repo {
	if c.dirRepo == nil {
		c.dirRepo = lo.Must(dirRepo.New(c.ctx, c.cfg.Storage.RootDirs, c.osAdapter))
	}

	return c.dirRepo
}
