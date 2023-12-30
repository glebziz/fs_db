package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDb_Set(t *testing.T) {
	_db := newTestDb(t)

	testGoN(t, testNumThread, func(t testing.TB) {
		for i := 0; i < testN; i++ {
			err := _db.Set(testCtx, gofakeit.UUID(), testContent)
			require.NoError(t, err)
		}
	})
}
