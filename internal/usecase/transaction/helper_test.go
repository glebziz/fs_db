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
	cleaner *mock_transaction.Mockcleaner
	fRepo   *mock_transaction.MockfileRepository
	txRepo  *mock_transaction.MocktxRepository

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
		cleaner: mock_transaction.NewMockcleaner(ctrl),
		fRepo:   mock_transaction.NewMockfileRepository(ctrl),
		txRepo:  mock_transaction.NewMocktxRepository(ctrl),
		idGen:   idGen,
	}
}

func (td *testDeps) newUseCase() *UseCase {
	return New(td.cleaner, td.fRepo, td.txRepo, td.idGen)
}
