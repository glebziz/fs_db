package dir

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/dir/mocks"
)

var (
	testMaxCount  = uint64(10_000)
	testName      = gofakeit.UUID()
	testName2     = gofakeit.UUID()
	testRootPath  = gofakeit.UUID()
	testRootPath2 = gofakeit.UUID()
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	dRepo   *mock_dir.MockdirRepository
	nameGen *mock_dir.Mockgenerator
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	idGen := mock_dir.NewMockgenerator(ctrl)
	idGen.EXPECT().
		Generate().
		AnyTimes().
		Return(testName)

	return &testDeps{
		dRepo:   mock_dir.NewMockdirRepository(ctrl),
		nameGen: idGen,
	}
}

func (d *testDeps) newUseCase() *useCase {
	return New(testMaxCount, d.dRepo, d.nameGen)
}
