package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Core(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing core",
			prepare: func(c *Container) {
				c.Core()
			},
		},
		{
			name:    "without existing cleaner",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.Core()
			require.NotNil(t, cl)
		})
	}
}
