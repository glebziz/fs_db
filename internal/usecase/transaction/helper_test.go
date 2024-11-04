package transaction

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/usecase/transaction/mocks"
)

const (
	testId       = "testId"
	testIsoLevel = fs_db.IsoLevelDefault
)

var (
	testCtx = model.StoreTxId(context.Background(), testId)
)

type prepareFunc func(td *testDeps)

type testDeps struct {
	fRepo  *mock_transaction.MockfileRepository
	txRepo *mock_transaction.MocktxRepository

	idGen *mock_transaction.Mockgenerator
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	idGen := mock_transaction.NewMockgenerator(ctrl)
	idGen.EXPECT().
		Generate().
		AnyTimes().
		Return(testId)

	return &testDeps{
		fRepo:  mock_transaction.NewMockfileRepository(ctrl),
		txRepo: mock_transaction.NewMocktxRepository(ctrl),
		idGen:  idGen,
	}
}

func (d *testDeps) newUseCase() *useCase {
	return New(d.fRepo, d.txRepo, d.idGen)
}
