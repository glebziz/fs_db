package di

import (
	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
	"github.com/glebziz/fs_db/internal/utils/must"
)

func (c *Container) DirRepo() *dirRepo.Repo {
	if c.dirRepo == nil {
		c.dirRepo = must.Must(dirRepo.New(c.ctx, c.cfg.Storage.RootDirs, c.osAdapter))
	}

	return c.dirRepo
}
