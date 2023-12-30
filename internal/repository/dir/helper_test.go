package dir

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
)

func testCreateDir(ctx context.Context, t *testing.T, dir *model.Dir) {
	t.Helper()

	_, err := manager.DB(ctx).Exec(ctx, `
		insert into dir(id, parent_path) VALUES ($1, $2)`,
		dir.Id, dir.ParentPath)
	require.NoError(t, err)
}

func testCreateFile(ctx context.Context, t *testing.T, file *model.File) {
	t.Helper()

	_, err := manager.DB(ctx).Exec(ctx, `
		insert into file(id, key, parent_path) VALUES ($1, $2, $3)`,
		file.Id, file.Key, file.ParentPath)
	require.NoError(t, err)
}

func testGetDir(ctx context.Context, t *testing.T, id string) *model.Dir {
	t.Helper()

	var d model.Dir
	rows, err := manager.DB(ctx).Query(ctx, `
		select id, parent_path 
		from dir where id = $1`, id)
	require.NoError(t, err)
	defer rows.Close()
	require.True(t, rows.Next())

	err = rows.Scan(&d.Id, &d.ParentPath)
	require.NoError(t, err)
	require.NoError(t, rows.Err())

	return &d
}
