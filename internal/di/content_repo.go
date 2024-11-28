package di

import (
	contentRepo "github.com/glebziz/fs_db/internal/repository/content"
)

func (c *Container) ContentRepo() *contentRepo.Repo {
	if c.contentRepo == nil {
		c.contentRepo = contentRepo.New()
	}

	return c.contentRepo
}
