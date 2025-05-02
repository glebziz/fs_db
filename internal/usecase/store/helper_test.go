package store

import (
	"context"
	"math/rand/v2"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/usecase/store/mocks"
)

var (
	testTxId  = gofakeit.UUID()
	testTxSeq = sequence.Next()

	testKey        = "testKey"
	testKey2       = "testKey2"
	testContentId  = gofakeit.UUID()
	testContentId2 = gofakeit.UUID()

	testContent = gofakeit.UUID()
	testSize    = uint64(10)
	testSize2   = uint64(8)
	testSize3   = uint64(1)
	testSize4   = uint64(9)

	testDirName  = gofakeit.UUID()
	testDirName2 = gofakeit.UUID()
	testDirName3 = gofakeit.UUID()
	testDirName4 = gofakeit.UUID()
	testRootPath = gofakeit.UUID()

	testCtx = model.StoreTxId(context.Background(), testTxId)
)

type prepareFunc func(td *testDeps) error

type randSource struct{}

func (randSource) Uint64() uint64 {
	return 4
}

type testDeps struct {
	dir *mock_store.MockdirUsecase

	cRepo   *mock_store.MockcontentRepository
	cfRepo  *mock_store.MockcontentFileRepository
	fRepo   *mock_store.MockfileRepository
	txRepo  *mock_store.MocktxRepository
	cWriter *mock_store.MockcontentWriter

	idGen  *mock_store.Mockgenerator
	reader *mock_store.MockReadSeekCloser
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	idGen := mock_store.NewMockgenerator(ctrl)
	idGen.EXPECT().
		Generate().
		AnyTimes().
		Return(testContentId)

	return &testDeps{
		dir:     mock_store.NewMockdirUsecase(ctrl),
		cRepo:   mock_store.NewMockcontentRepository(ctrl),
		cfRepo:  mock_store.NewMockcontentFileRepository(ctrl),
		fRepo:   mock_store.NewMockfileRepository(ctrl),
		txRepo:  mock_store.NewMocktxRepository(ctrl),
		cWriter: mock_store.NewMockcontentWriter(ctrl),
		idGen:   idGen,
		reader:  mock_store.NewMockReadSeekCloser(ctrl),
	}
}

func (d *testDeps) newUseCase() *UseCase {
	return New(
		d.dir, d.cRepo,
		d.cfRepo, d.fRepo,
		d.txRepo, d.cWriter,
		d.idGen, rand.New(randSource{}),
	)
}
