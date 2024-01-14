package transaction

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

var (
	testCtx = context.Background()
)

func testCreateTransaction(t *testing.T, repo *rep, tx model.Transaction) {
	t.Helper()

	repo.storage.Store(tx.Id, &tx)
}

func testGetTransaction(t *testing.T, repo *rep, id string) *model.Transaction {
	t.Helper()

	tx, ok := repo.storage.Load(id)
	require.True(t, ok)

	return tx
}
