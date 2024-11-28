package di

import (
	"github.com/glebziz/fs_db/internal/utils/generator"
)

func (c *Container) Gen() *generator.Gen {
	if c.gen == nil {
		c.gen = generator.New()
	}

	return c.gen
}
