package store

import (
	"context"
	"io"
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
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
	testSize2   = uint64(8)
	testSize3   = uint64(1)
	testSize4   = uint64(9)
	testReader  = io.NopCloser(strings.NewReader(testContent))

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

type closer struct {
	io.Reader
	count int
}

func (c *closer) Close() error {
	c.count++
	return nil
}

func testNewCloser(t *testing.T, r io.Reader, times int) io.ReadCloser {
	t.Helper()

	c := closer{
		Reader: r,
	}

	t.Cleanup(func() {
		require.Equal(t, times, c.count)
	})

	return &c
}

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
		rand.New(randSource{}),
	)
}
