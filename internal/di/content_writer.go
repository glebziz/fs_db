package di

import (
	"github.com/glebziz/fs_db/internal/usecase/writer"
)

func (c *Container) ContentWriter() *writer.UseCase {
	if c.contentWriter == nil {
		c.contentWriter = writer.New(c.ContentRepo())
	}

	return c.contentWriter
}
