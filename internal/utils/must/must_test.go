package must

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {
	t.Parallel()

	const (
		mustValue int64 = iota + 1
	)

	t.Run("success", func(t *testing.T) {
		val := Must(mustValue, nil)
		require.Equal(t, mustValue, val)
	})

	t.Run("error", func(t *testing.T) {
		defer func() {
			err, ok := recover().(error)
			if !ok {
				return
			}

			require.ErrorIs(t, err, assert.AnError)
		}()

		_ = Must(mustValue, assert.AnError)
	})
}
