package content

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func testCreateFile(t *testing.T, path string, content []byte) {
	t.Helper()

	err := os.WriteFile(path, content, 0666)
	require.NoError(t, err)
}
