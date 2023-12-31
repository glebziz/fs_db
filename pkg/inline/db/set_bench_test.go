package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func BenchmarkDb_Set(b *testing.B) {
	_db := newTestDb(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		key := gofakeit.UUID()
		b.StartTimer()

		err := _db.Set(testCtx, key, testContent)
		require.NoError(b, err)
	}
}
