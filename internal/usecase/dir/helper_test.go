package dir

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/dir/mocks"
)

const (
	testMaxCount  = uint64(10_000)
	testName      = "testName"
	testName2     = "testName2"
	testNewName   = "testNewName"
	testRootPath  = "testRootPath"
	testRootPath2 = "testRootPath2"
)

type prepareFunc func(td *testDeps)

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
		Return(testNewName)

	return &testDeps{
		dRepo:   mock_dir.NewMockdirRepository(ctrl),
		nameGen: idGen,
	}
}

func (d *testDeps) newUseCase() *UseCase {
	return New(testMaxCount, d.dRepo, d.nameGen)
}
