package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Badger(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing manager",
			prepare: func(c *Container) {
				c.Badger()
			},
		},
		{
			name:    "without existing manager",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			m := c.Badger()
			require.NotNil(t, m)
		})
	}
}
