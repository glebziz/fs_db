package di

import (
	"math/rand/v2"
)

func (c *Container) Rand() *rand.Rand {
	if c.rand == nil {
		c.rand = rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))
	}

	return c.rand
}
