package db

import (
	"bytes"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func BenchmarkDb_SetReader(b *testing.B) {
	_db := newTestDb(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		key := gofakeit.UUID()
		reader := bytes.NewReader(testContent)
		b.StartTimer()

		err := _db.SetReader(testCtx, key, reader, testSize)
		require.NoError(b, err)
	}
}
