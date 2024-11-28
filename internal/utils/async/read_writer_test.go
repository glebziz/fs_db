package async

import (
	"errors"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadWriter(t *testing.T) {
	const (
		data = "some content"
	)

	t.Run("success read and write", func(t *testing.T) {
		t.Parallel()

		var (
			wg sync.WaitGroup
			rw = NewReadWriter()
		)

		wg.Add(1)
		go func() {
			p := make([]byte, len(data))

			wg.Done()

			n, err := rw.Read(p)
			require.Len(t, p, n)
			require.NoError(t, err)
		}()

		wg.Wait()
		n, err := rw.Write([]byte(data))
		require.Len(t, data, n)
		require.NoError(t, err)
	})

	t.Run("success read after closing", func(t *testing.T) {
		t.Parallel()

		var (
			rw = NewReadWriter()
		)

		n, err := rw.Write([]byte(data))
		require.Len(t, data, n)
		require.NoError(t, err)

		err = rw.Close()
		require.NoError(t, err)

		d, err := io.ReadAll(rw)
		require.NoError(t, err)
		require.EqualValues(t, data, d)
	})

	t.Run("write after error", func(t *testing.T) {
		t.Parallel()

		var (
			rw = NewReadWriter()
		)

		rw.SetError(assert.AnError)

		n, err := rw.Write([]byte(data))
		require.Zero(t, n)
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("set multi errors", func(t *testing.T) {
		t.Parallel()

		var (
			rw = NewReadWriter()

			secondErr = errors.New("second")
		)

		rw.SetError(assert.AnError)
		rw.SetError(secondErr)

		err := rw.checkErr()
		require.ErrorIs(t, err, assert.AnError)
		require.ErrorIs(t, err, secondErr)
	})
}
