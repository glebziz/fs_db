package db

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func BenchmarkDb_Create(b *testing.B) {
	_db := newTestDb(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		key := gofakeit.UUID()
		b.StartTimer()

		f, err := _db.Create(testCtx, key)
		require.NoError(b, err)

		n, err := f.Write(testContent)
		require.Len(b, testContent, n)
		require.NoError(b, err)

		err = f.Close()
		require.NoError(b, err)
	}
}
