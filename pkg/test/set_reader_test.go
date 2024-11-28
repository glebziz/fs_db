package db

import (
	"bytes"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDb_SetReader(t *testing.T) {
	t.Parallel()

	_db := newTestDb(t)

	testGoN(t, testNumThread, func(t testing.TB) {
		for i := 0; i < testN; i++ {
			err := _db.SetReader(testCtx, gofakeit.UUID(), bytes.NewReader(testContent))
			require.NoError(t, err)
		}
	})
}
