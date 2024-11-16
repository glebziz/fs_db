package file

import (
	"context"
	"fmt"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model/transactor"
	file_mock "github.com/glebziz/fs_db/internal/repository/file/mocks"
)

const (
	testKey        = "testKey"
	testKey2       = "testKey2"
	testTxId       = "00010203-0405-0607-0809-0a0b0c0d0e0f"
	testTxId2      = "0f0e0d0c-0c0a-0908-0706-050403020100"
	testContentId  = "00000101-0202-0303-0404-050506060707"
	testContentId2 = "08080909-0a0a-0b0b-0c0c-0d0d0e0e0f0f"
	testSeq        = 1
	testSeq2       = 2
)

type testDeps struct {
	p  badger.Provider
	qm *file_mock.MockQueryManager
}

type prepareFunc func(td *testDeps)
type prepareIntFunc func(t *testing.T, p badger.Provider)

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)
	p := file_mock.NewMockProvider(ctrl)
	qm := file_mock.NewMockQueryManager(ctrl)

	p.EXPECT().
		DB(gomock.Any()).
		MaxTimes(1).
		Return(qm)

	p.EXPECT().
		RunTransaction(gomock.Any(), gomock.Any()).
		MaxTimes(1).
		DoAndReturn(func(ctx context.Context, fn transactor.TransactionFn) error {
			return fn(ctx)
		})

	return &testDeps{
		p:  p,
		qm: qm,
	}
}

func (td *testDeps) newRep() *rep {
	return New(td.p)
}

func newTestRep(t *testing.T) *rep {
	t.Helper()

	var (
		dbPath = path.Join(t.TempDir(), fmt.Sprintf("test_file_%s", gofakeit.UUID()))
	)

	manager, err := badger.New(dbPath)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = manager.Close()
		require.NoError(t, err)
	})

	return New(manager)
}
