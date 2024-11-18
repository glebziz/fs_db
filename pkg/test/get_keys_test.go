package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestDb_GetKeys(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	var (
		key  = gofakeit.UUID()
		key2 = gofakeit.UUID()
		key3 = gofakeit.UUID()
	)

	err := _db.Set(testCtx, key, testContent)
	require.NoError(t, err)

	err = _db.Set(testCtx, key2, testContent)
	require.NoError(t, err)

	err = _db.Set(testCtx, key3, testContent)
	require.NoError(t, err)

	testGoN(t, testNumThread, func(t testing.TB) {
		var keys []string
		keys, err = _db.GetKeys(testCtx)
		require.NoError(t, err)
		require.True(t, gomock.InAnyOrder(keys).Matches([]string{key, key2, key3}))
	})
}
