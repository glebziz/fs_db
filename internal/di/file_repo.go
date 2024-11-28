package di

import (
	fileRepo "github.com/glebziz/fs_db/internal/repository/file"
)

func (c *Container) FileRepo() *fileRepo.Repo {
	if c.fileRepo == nil {
		c.fileRepo = fileRepo.New(c.Badger())
	}

	return c.fileRepo
}
