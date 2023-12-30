package root

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/root/mocks"
)

var (
	testRootPath   = gofakeit.UUID()
	testRootPath2  = gofakeit.UUID()
	testRootFree   = uint64(1000)
	testRootFree2  = uint64(100)
	testRootCount  = uint64(1)
	testRootCount2 = uint64(2)
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	manager *mock_root.MockdiskManager
	repo    *mock_root.MockdirRepository
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		manager: mock_root.NewMockdiskManager(ctrl),
		repo:    mock_root.NewMockdirRepository(ctrl),
	}
}

func (d *testDeps) newUseCase(rootDirs []string) *useCase {
	return New(rootDirs, d.manager, d.repo)
}
