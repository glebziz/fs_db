package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDb_Get(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	key := gofakeit.UUID()
	err := _db.Set(testCtx, key, testContent)
	require.NoError(t, err)

	testGoN(t, testNumThread, func(t testing.TB) {
		c, err := _db.Get(testCtx, key)
		require.NoError(t, err)
		require.Equal(t, testContent, c)
	})
}
