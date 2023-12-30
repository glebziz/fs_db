package dir

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/dir/mocks"
)

var (
	testMaxCount  = uint64(10_000)
	testId        = gofakeit.UUID()
	testRootPath  = gofakeit.UUID()
	testRootPath2 = gofakeit.UUID()
	testFileCount = uint64(1000)
	testSize      = uint64(10)
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	root  *mock_dir.MockrootUseCase
	dRepo *mock_dir.MockdirRepository
	idGen *mock_dir.Mockgenerator
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	idGen := mock_dir.NewMockgenerator(ctrl)
	idGen.EXPECT().
		Generate().
		AnyTimes().
		Return(testId)

	return &testDeps{
		root:  mock_dir.NewMockrootUseCase(ctrl),
		dRepo: mock_dir.NewMockdirRepository(ctrl),
		idGen: idGen,
	}
}

func (d *testDeps) newUseCase() *useCase {
	return New(testMaxCount, d.root, d.dRepo, d.idGen)
}
