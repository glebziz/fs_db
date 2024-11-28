package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func BenchmarkDb_GetKeys(b *testing.B) {
	_db := newTestDb(b)
	keys := make([]string, 0, b.N)
	for i := 0; i < 100; i++ {
		key := gofakeit.UUID()

		err := _db.Set(testCtx, key, testContent)
		require.NoError(b, err)

		keys = append(keys, key)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actKeys, err := _db.GetKeys(testCtx)
		require.NoError(b, err)
		require.Equal(b, len(keys), len(actKeys))
	}
}
