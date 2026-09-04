package db

import (
	"bytes"
	"context"
	"sync"
	"testing"
)

const (
	testN         = 100
	testNumThread = 10
)

var (
	testCtx     = context.Background()
	testContent = bytes.Repeat([]byte("1"), 1<<15)
)

func testGoN(t testing.TB, n int, fn func(t testing.TB)) {
	t.Helper()

	wg := sync.WaitGroup{}
	for range n {

		wg.Go(func() {
			fn(t)
		})
	}

	wg.Wait()
}
