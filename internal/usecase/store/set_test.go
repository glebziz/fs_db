package store

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	pkgModel "github.com/glebziz/fs_db/pkg/model"
)

func TestUseCase_Set_Success(t *testing.T) {
	t.Parallel()

	var (
		dir = model.Dir{
			Id:         testDirId,
			ParentPath: testRootPath,
		}
		file = model.File{
			Id:         testId,
			Key:        testKey,
			ParentPath: dir.GetPath(),
		}
		content = model.Content{
			Size:   testSize,
			Reader: testReader,
		}
	)
	td := newTestDeps(t)

	td.dir.EXPECT().
		Select(gomock.Any(), testSize).
		Return(&dir, nil)

	td.fRepo.EXPECT().
		Store(gomock.Any(), file).
		Return(nil)

	td.cRepo.EXPECT().
		Store(gomock.Any(), file.GetPath(), &content).
		DoAndReturn(func(_ context.Context, _ string, content *model.Content) error {
			data, err := io.ReadAll(content.Reader)
			require.NoError(t, err)
			require.Equal(t, testContent, string(data))

			return nil
		})

	uc := td.newUseCase()

	err := uc.Set(context.Background(), testKey, &content)

	require.NoError(t, err)
}

func TestUseCase_Set_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		prepare prepareFunc
	}{
		{
			name: "empty key",
			prepare: func(td *testDeps) error {
				return pkgModel.EmptyKeyErr
			},
		},
		{
			name: "select dir",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Select(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "file store",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Select(gomock.Any(), gomock.Any()).
					Return(&model.Dir{
						Id:         testDirId,
						ParentPath: testRootPath,
					}, nil)

				td.fRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content store",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Select(gomock.Any(), gomock.Any()).
					Return(&model.Dir{
						Id:         testDirId,
						ParentPath: testRootPath,
					}, nil)

				td.fRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil)

				td.cRepo.EXPECT().
					Store(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				content = model.Content{
					Size:   testSize,
					Reader: testReader,
				}
			)
			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			err := uc.Set(context.Background(), tc.key, &content)

			require.ErrorIs(t, err, wantErr)
		})
	}
}
