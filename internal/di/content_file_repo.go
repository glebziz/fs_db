package di

import (
	contentFileRepo "github.com/glebziz/fs_db/internal/repository/content_file"
)

func (c *Container) ContentFileRepo() *contentFileRepo.Repo {
	if c.contentFileRepo == nil {
		c.contentFileRepo = contentFileRepo.New(c.Badger())
	}

	return c.contentFileRepo
}
