package di

import (
	"github.com/samber/lo"

	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
)

func (c *Container) DirRepo() *dirRepo.Repo {
	if c.dirRepo == nil {
		c.dirRepo = lo.Must(dirRepo.New(c.cfg.Storage.RootDirs))
	}

	return c.dirRepo
}
