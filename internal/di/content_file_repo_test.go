package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_ContentFileRepo(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing repo",
			prepare: func(c *Container) {
				c.ContentFileRepo()
			},
		},
		{
			name:    "without existing repo",
			prepare: func(c *Container) {},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := newContainer(t)
			tc.prepare(c)

			cl := c.ContentFileRepo()
			require.NotNil(t, cl)
		})
	}
}
