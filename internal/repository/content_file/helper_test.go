package file

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/db"
	"github.com/glebziz/fs_db/internal/model"
)

func newTestRep(t *testing.T) (*rep, context.Context) {
	t.Helper()

	var (
		dbPath = path.Join(os.TempDir(), fmt.Sprintf("test_dir_%s.db", gofakeit.UUID()))
	)

	manager, err := db.New(context.Background(), dbPath)
	require.NoError(t, err)
	t.Cleanup(func() {
		manager.Close()
		err = os.Remove(dbPath)
		require.NoError(t, err)
	})

	return New(manager), context.Background()
}

func testCreateContentFile(ctx context.Context, t *testing.T, p db.Provider, file *model.ContentFile) {
	t.Helper()

	_, err := p.DB(ctx).Exec(ctx, `
		insert into content_file(id, parent_path) VALUES ($1, $2)`,
		file.Id, file.Parent)
	require.NoError(t, err)
}

func testGetContentFile(ctx context.Context, t *testing.T, p db.Provider, id string) *model.ContentFile {
	t.Helper()

	var file model.ContentFile
	rows, err := p.DB(ctx).Query(ctx, `
		select id, parent_path 
		from content_file where id = $1`, id)
	require.NoError(t, err)
	defer rows.Close()
	require.True(t, rows.Next())

	err = rows.Scan(&file.Id, &file.Parent)
	require.NoError(t, err)
	require.NoError(t, rows.Err())

	return &file
}
