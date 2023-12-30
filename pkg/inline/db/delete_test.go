package db

import (
	"sync/atomic"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/pkg/model"
)

func TestDb_Delete(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	var (
		count = int32(0)
		key   = gofakeit.UUID()
	)

	err := _db.Set(testCtx, key, testContent)
	require.NoError(t, err)

	c, err := _db.Get(testCtx, key)
	require.NoError(t, err)
	require.Equal(t, testContent, c)

	testGoN(t, testNumThread, func(t testing.TB) {
		err = _db.Delete(testCtx, key)

		if err == nil {
			atomic.AddInt32(&count, 1)
		} else {
			require.ErrorIs(t, err, model.NotFoundErr)
		}
	})

	require.Equal(t, int32(1), count)

	_, err = _db.Get(testCtx, key)
	require.ErrorIs(t, err, model.NotFoundErr)
}
