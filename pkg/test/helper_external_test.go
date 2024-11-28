//go:build external

package db

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/app"
	externalDb "github.com/glebziz/fs_db/pkg/external/db"
)

func newTestDb(t testing.TB) fs_db.DB {
	t.Helper()

	var (
		wg sync.WaitGroup

		port = 1000 + rand.Intn(100)
	)

	dir, err := os.MkdirTemp("", "fs_db_test")
	require.NoError(t, err)

	err = os.Chmod(dir, 0750)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	a, err := app.New(ctx, config.Config{
		Port: port,
		Storage: config.Storage{
			MaxDirCount: 100,
			DbPath:      path.Join(dir, "test_db"),
			RootDirs:    []string{path.Join(dir, gofakeit.UUID()), path.Join(dir, gofakeit.UUID())},
			GCPeriod:    time.Minute,
		},
		WPool: config.WPool{
			NumWorkers:   runtime.GOMAXPROCS(0),
			SendDuration: time.Millisecond,
		},
	})
	require.NoError(t, err)

	wg.Add(1)
	go func() {
		defer wg.Done()

		runErr := a.Run(ctx)
		require.NoError(t, runErr)
	}()

	t.Cleanup(func() {
		cancel()

		wg.Wait()

		stopErr := a.Stop()
		require.NoError(t, stopErr)
	})

	_db, err := externalDb.New(testCtx, fmt.Sprintf("localhost:%d", port))
	require.NoError(t, err)

	return _db
}
