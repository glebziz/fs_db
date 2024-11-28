package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDb_Create(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	testGoN(t, testNumThread, func(t testing.TB) {
		for i := 0; i < testN; i++ {
			f, err := _db.Create(testCtx, gofakeit.UUID())
			require.NoError(t, err)

			n, err := f.Write(testContent)
			require.Len(t, testContent, n)
			require.NoError(t, err)

			err = f.Close()
			require.NoError(t, err)
		}
	})
}
