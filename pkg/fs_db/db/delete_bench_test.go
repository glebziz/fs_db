package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func BenchmarkDb_Delete(b *testing.B) {
	_db := newTestDb(b)
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i++ {
		key := gofakeit.UUID()

		err := _db.Set(testCtx, key, testContent)
		require.NoError(b, err)

		keys = append(keys, key)
	}

	b.ResetTimer()

	for _, key := range keys {
		err := _db.Delete(testCtx, key)
		require.NoError(b, err)
	}
}
