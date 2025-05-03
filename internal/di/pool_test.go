package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Pool(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing pool",
			prepare: func(c *Container) {
				c.Pool()
			},
		},
		{
			name:    "without existing pool",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.Pool()
			require.NotNil(t, cl)
		})
	}
}
