package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Cleaner(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing cleaner",
			prepare: func(c *Container) {
				c.Cleaner()
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

			cl := c.Cleaner()
			require.NotNil(t, cl)
		})
	}
}
