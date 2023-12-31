package dir

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func testCreateFile(ctx context.Context, t *testing.T, file *model.File) {
	t.Helper()

	_, err := manager.DB(ctx).Exec(ctx, `
		insert into file(id, key, parent_path) VALUES ($1, $2, $3)`,
		file.Id, file.Key, file.ParentPath)
	require.NoError(t, err)
}

func testGetFile(ctx context.Context, t *testing.T, key string) *model.File {
	t.Helper()

	var file model.File
	rows, err := manager.DB(ctx).Query(ctx, `
		select id, key, parent_path 
		from file where key = $1`, key)
	require.NoError(t, err)
	defer rows.Close()
	require.True(t, rows.Next())

	err = rows.Scan(&file.Id, &file.Key, &file.ParentPath)
	require.NoError(t, err)
	require.NoError(t, rows.Err())

	return &file
}
