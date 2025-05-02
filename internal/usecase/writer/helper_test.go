package writer

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/writer/mocks"
)

type prepareFunc func(td *testDeps)

type testDeps struct {
	t    *testing.T
	ctrl *gomock.Controller

	cRepo *mocks.MockcontentRepository
}

func newTestDeps(t *testing.T) *testDeps {
	t.Helper()

	ctrl := gomock.NewController(t)
	return &testDeps{
		t:     t,
		ctrl:  ctrl,
		cRepo: mocks.NewMockcontentRepository(ctrl),
	}
}

func (d *testDeps) newUseCase() *UseCase {
	return New(d.cRepo)
}
