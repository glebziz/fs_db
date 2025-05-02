package db

import (
	"io"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
)

func TestDb_Open(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	key := gofakeit.UUID()

	{
		r, err := _db.Open(testCtx, key)
		require.ErrorIs(t, err, fs_db.ErrNotFound)
		require.Nil(t, r)

		err = _db.Set(testCtx, key, testContent)
		require.NoError(t, err)
	}

	testGoN(t, testNumThread, func(t testing.TB) {
		f, err := _db.Open(testCtx, key)
		require.NoError(t, err)

		pos, err := f.Seek(int64(len(testContent)/2), io.SeekStart)
		require.NoError(t, err)
		require.EqualValues(t, len(testContent)/2, pos)

		c, err := io.ReadAll(f)
		require.NoError(t, err)
		require.Equal(t, testContent[pos:], c)

		f.Close()
	})
}
