package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Dir(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing dir usecase",
			prepare: func(c *Container) {
				c.Dir()
			},
		},
		{
			name:    "without existing dir usecase",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.Dir()
			require.NotNil(t, cl)
		})
	}
}
