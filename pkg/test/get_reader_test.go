package db

import (
	"io"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
)

func TestDb_GetReader(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	key := gofakeit.UUID()

	r, err := _db.GetReader(testCtx, key)
	require.ErrorIs(t, err, fs_db.ErrNotFound)
	require.Nil(t, r)

	err = _db.Set(testCtx, key, testContent)
	require.NoError(t, err)

	testGoN(t, testNumThread, func(t testing.TB) {
		r, err := _db.GetReader(testCtx, key)
		require.NoError(t, err)

		c, err := io.ReadAll(r)
		require.NoError(t, err)
		require.Equal(t, testContent, c)

		r.Close()
	})
}
