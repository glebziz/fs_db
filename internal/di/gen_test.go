package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Gen(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing generator",
			prepare: func(c *Container) {
				c.Gen()
			},
		},
		{
			name:    "without existing generator",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.Gen()
			require.NotNil(t, cl)
		})
	}
}
