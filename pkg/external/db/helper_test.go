package db

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/glebziz/fs_db/config"
	storeService "github.com/glebziz/fs_db/internal/delivery/grpc/store"
	store "github.com/glebziz/fs_db/internal/proto"
	inlineDb "github.com/glebziz/fs_db/pkg/inline/db"
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

func newTestDb(t testing.TB) *db {
	t.Helper()

	var (
		cancel = make(chan struct{})
		runWg  = sync.WaitGroup{}
		doneWg = sync.WaitGroup{}
		port   = 1000 + rand.Intn(100)
	)

	dir, err := os.MkdirTemp("", "fs_db_test")
	require.NoError(t, err)

	cl, err := inlineDb.New(testCtx, &config.Storage{
		MaxDirCount: 100,
		DbPath:      path.Join(dir, "test.db"),
		RootDirs:    []string{path.Join(dir, gofakeit.UUID()), path.Join(dir, gofakeit.UUID())},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		close(cancel)
		doneWg.Wait()

		err = cl.Close()
		require.NoError(t, err)

		err = os.RemoveAll(dir)
		require.NoError(t, err)
	})

	runWg.Add(1)
	doneWg.Add(1)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		require.NoError(t, err)

		s := grpc.NewServer()
		store.RegisterStoreV1Server(s, storeService.New(cl.GetUseCase()))
		go func() {
			runWg.Done()
			s.Serve(lis)
		}()

		select {
		case <-cancel:
			s.GracefulStop()
			doneWg.Done()
		}
	}()

	runWg.Wait()

	_db, err := New(testCtx, fmt.Sprintf("localhost:%d", port))
	require.NoError(t, err)

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
