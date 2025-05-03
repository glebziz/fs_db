package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Rand(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing randomizer",
			prepare: func(c *Container) {
				c.Rand()
			},
		},
		{
			name:    "without existing randomizer",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.Rand()
			require.NotNil(t, cl)
		})
	}
}
