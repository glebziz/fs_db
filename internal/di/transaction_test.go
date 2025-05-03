package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer_Transaction(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "with existing usecase",
			prepare: func(c *Container) {
				c.Transaction()
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

			cl := c.Transaction()
			require.NotNil(t, cl)
		})
	}
}
