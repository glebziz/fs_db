package db

import (
	"bytes"
	"context"
	"os"
	"path"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/config"
)

const (
	testN         = 100
	testNumThread = 10
)

var (
	testCtx     = context.Background()
	testContent = bytes.Repeat([]byte("1"), 1<<15)
	testSize    = uint64(len(testContent))
)

func newTestDb(t testing.TB) fs_db.DB {
	t.Helper()

	dir, err := os.MkdirTemp("", "fs_db_test")
	require.NoError(t, err)

	_db, err := New(testCtx, &config.Storage{
		MaxDirCount: 100,
		DbPath:      path.Join(dir, "test.db"),
		RootDirs:    []string{path.Join(dir, gofakeit.UUID()), path.Join(dir, gofakeit.UUID())},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = _db.Close()
		require.NoError(t, err)

		err = os.RemoveAll(dir)
		require.NoError(t, err)
	})

	return _db
}

func testGoN(t testing.TB, n int, fn func(t testing.TB)) {
	t.Helper()

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fn(t)
		}()
	}

	wg.Wait()
}
