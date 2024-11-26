package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/core"
	mock_core "github.com/glebziz/fs_db/internal/usecase/core/mocks"
)

const (
	testTxId  = "testTxId"
	testTxId2 = "testTxId2"
	testTxId3 = "testTxId3"

	testKey  = "testKey"
	testKey2 = "testKey2"
	testKey3 = "testKey3"

	testContentId  = "testContentId"
	testContentId2 = "testContentId2"
	testContentId3 = "testContentId3"
	testContentId4 = "testContentId4"
	testContentId5 = "testContentId5"
	testContentId6 = "testContentId6"
)

type initUseCaseFunc func(td *testDeps) (*UseCase, model.FileFilter)
type requireUseCaseFunc func(t *testing.T, u *UseCase)

type testDeps struct {
	t        *testing.T
	fileRepo *mock_core.MockfileRepository
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		t:        t,
		fileRepo: mock_core.NewMockfileRepository(ctrl),
	}
}

func (td *testDeps) newUseCase() *UseCase {
	return New(td.fileRepo)
}

func requireEqualFiles(t *testing.T, a, b []model.File) {
	t.Helper()

	require.Len(t, b, len(a), fmt.Sprintf("%v\n%v", a, b))
	for i := range b {
		b[i].Seq = 0
	}

	require.True(t, gomock.InAnyOrder(a).Matches(b), fmt.Sprintf("%v\n%v", a, b))
}

func requireEqualFile(t *testing.T, a, b model.File) {
	t.Helper()

	b.Seq = 0
	require.Equal(t, a, b)
}

func (u *UseCase) testStore(td *testDeps, f model.File) {
	td.t.Helper()

	tx, ok := u.txStore.Get(f.TxId)
	if !ok {
		tx = &core.Transaction{}
		u.txStore.Put(f.TxId, tx)
	}

	u.storeToTx(tx, f)
}

func (u *UseCase) testAddEmptyTx(td *testDeps, txId string, keys ...string) {
	td.t.Helper()

	tx := &core.Transaction{}
	u.txStore.Put(txId, tx)
	for _, key := range keys {
		tx.PushBack((&core.Node[model.File]{}).SetV(model.File{
			Key:  key,
			TxId: txId,
		}))
		tx.File(key).PopBack()
	}
}
