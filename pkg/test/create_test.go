package db

import (
	"io"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDb_Create(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	const (
		offset int64 = iota
	)

	testGoN(t, testNumThread, func(t testing.TB) {
		for i := 0; i < testN; i++ {
			key := gofakeit.UUID()

			f, err := _db.Create(testCtx, key)
			require.NoError(t, err)

			pos, err := f.Seek(offset, io.SeekStart)
			require.NoError(t, err)
			require.Equal(t, offset, pos)

			n, err := f.Write(testContent)
			require.Len(t, testContent, n)
			require.NoError(t, err)

			err = f.Close()
			require.NoError(t, err)

			content, err := _db.Get(testCtx, key)
			require.NoError(t, err)
			require.Equal(t, append(make([]byte, offset), testContent...), content)
		}
	})
}
