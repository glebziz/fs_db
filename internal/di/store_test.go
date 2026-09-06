package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Store(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing usecase",
			prepare: func(c *Container) {
				c.Store()
			},
		},
		{
			name:    "without existing usecase",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.Store()
			require.NotNil(t, cl)
		})
	}
}
