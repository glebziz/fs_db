package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_ContentWriter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing writer",
			prepare: func(c *Container) {
				c.ContentWriter()
			},
		},
		{
			name:    "without existing writer",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.ContentWriter()
			require.NotNil(t, cl)
		})
	}
}
