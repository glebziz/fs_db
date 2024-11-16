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
	cf_mock "github.com/glebziz/fs_db/internal/repository/content_file/mocks"
)

const (
	testId     = "testId"
	testParent = "testParent"
)

type testDeps struct {
	p  badger.Provider
	qm *cf_mock.MockQueryManager
}

type prepareFunc func(td *testDeps)
type prepareIntFunc func(t *testing.T, p badger.Provider)

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)
	p := cf_mock.NewMockProvider(ctrl)
	qm := cf_mock.NewMockQueryManager(ctrl)

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
		dbPath = path.Join(t.TempDir(), fmt.Sprintf("test_content_file_%s", gofakeit.UUID()))
	)

	manager, err := badger.New(dbPath)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = manager.Close()
		require.NoError(t, err)
	})

	return New(manager)
}
