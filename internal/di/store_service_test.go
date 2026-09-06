package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_StoreService(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing service",
			prepare: func(c *Container) {
				c.StoreService()
			},
		},
		{
			name:    "without existing service",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.StoreService()
			require.NotNil(t, cl)
		})
	}
}
