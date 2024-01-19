package transaction

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/usecase/transaction/mocks"
)

var (
	testId       = gofakeit.UUID()
	testIsoLevel = fs_db.IsoLevelDefault

	testCtx = model.StoreTxId(context.Background(), testId)
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	cleaner *mock_transaction.Mockcleaner

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
		cleaner: mock_transaction.NewMockcleaner(ctrl),
		fRepo:   mock_transaction.NewMockfileRepository(ctrl),
		txRepo:  mock_transaction.NewMocktxRepository(ctrl),
		idGen:   idGen,
	}
}

func (d *testDeps) newUseCase() *useCase {
	return New(
		d.cleaner, d.fRepo,
		d.txRepo, d.idGen,
	)
}
