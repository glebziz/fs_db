package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
)

func TestDb_Delete(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	var (
		key = gofakeit.UUID()
	)

	err := _db.Set(testCtx, key, testContent)
	require.NoError(t, err)

	c, err := _db.Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, testContent, c)

	testGoN(t, testNumThread, func(t testing.TB) {
		err = _db.Delete(testCtx, key)
		require.NoError(t, err)
	})

	_, err = _db.Get(testCtx, key)
	require.ErrorIs(t, err, fs_db.NotFoundErr)
}
