package store

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/usecase/store/mocks"
)

var (
	testTxId = gofakeit.UUID()
	testTxTs = gofakeit.Date()

	testKey       = gofakeit.UUID()
	testContentId = gofakeit.UUID()

	testContent = gofakeit.UUID()
	testSize    = uint64(10)
	testReader  = io.NopCloser(strings.NewReader(testContent))

	testDirId    = gofakeit.UUID()
	testRootPath = gofakeit.UUID()

	testCtx = model.StoreTxId(context.Background(), testTxId)
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	dir *mock_store.MockdirUsecase

	cRepo  *mock_store.MockcontentRepository
	cfRepo *mock_store.MockcontentFileRepository
	fRepo  *mock_store.MockfileRepository
	txRepo *mock_store.MocktxRepository

	idGen *mock_store.Mockgenerator
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	idGen := mock_store.NewMockgenerator(ctrl)
	idGen.EXPECT().
		Generate().
		AnyTimes().
		Return(testContentId)

	return &testDeps{
		dir:    mock_store.NewMockdirUsecase(ctrl),
		cRepo:  mock_store.NewMockcontentRepository(ctrl),
		cfRepo: mock_store.NewMockcontentFileRepository(ctrl),
		fRepo:  mock_store.NewMockfileRepository(ctrl),
		txRepo: mock_store.NewMocktxRepository(ctrl),
		idGen:  idGen,
	}
}

func (d *testDeps) newUseCase() *useCase {
	return New(
		d.dir, d.cRepo,
		d.cfRepo, d.fRepo,
		d.txRepo, d.idGen,
	)
}
