package store

import (
	"io"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/store/mocks"
)

var (
	testId  = gofakeit.UUID()
	testKey = gofakeit.UUID()

	testContent = gofakeit.UUID()
	testSize    = uint64(10)
	testReader  = io.NopCloser(strings.NewReader(testContent))

	testDirId    = gofakeit.UUID()
	testRootPath = gofakeit.UUID()
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	dir *mock_store.MockdirUsecase

	cRepo *mock_store.MockcontentRepository
	fRepo *mock_store.MockfileRepository

	idGen *mock_store.Mockgenerator
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	idGen := mock_store.NewMockgenerator(ctrl)
	idGen.EXPECT().
		Generate().
		AnyTimes().
		Return(testId)

	return &testDeps{
		dir:   mock_store.NewMockdirUsecase(ctrl),
		cRepo: mock_store.NewMockcontentRepository(ctrl),
		fRepo: mock_store.NewMockfileRepository(ctrl),
		idGen: idGen,
	}
}

func (d *testDeps) newUseCase() *useCase {
	return New(d.dir, d.cRepo, d.fRepo, d.idGen)
}
