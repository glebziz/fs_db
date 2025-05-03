package content

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/repository/content/mocks"
)

type prepareFunc func(td *testDeps)

type testDeps struct {
	os *mocks.MockosAdapter
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		os: mocks.NewMockosAdapter(ctrl),
	}
}

func (td *testDeps) newRepo() *Repo {
	return New(td.os)
}

func testCreateFile(t *testing.T, path string, content []byte) {
	t.Helper()

	err := os.WriteFile(path, content, 0666)
	require.NoError(t, err)
}
