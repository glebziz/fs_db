package store

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Set_Success(t *testing.T) {
	t.Parallel()

	var (
		dir = model.Dir{
			Id:         testDirId,
			ParentPath: testRootPath,
		}
		cFile = model.ContentFile{
			Id:         testContentId,
			ParentPath: dir.GetPath(),
		}
		file = model.File{
			Key:       testKey,
			ContentId: testContentId,
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

	td.cRepo.EXPECT().
		Store(gomock.Any(), cFile.GetPath(), &content).
		DoAndReturn(func(_ context.Context, _ string, content *model.Content) error {
			data, err := io.ReadAll(content.Reader)
			require.NoError(t, err)
			require.Equal(t, testContent, string(data))

			return nil
		})

	td.cfRepo.EXPECT().
		Store(gomock.Any(), cFile).
		Return(nil)

	td.fRepo.EXPECT().
		Store(gomock.Any(), testTxId, file).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Set(testCtx, testKey, &content)

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
				return fs_db.EmptyKeyErr
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
			name: "content store",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Select(gomock.Any(), gomock.Any()).
					Return(&model.Dir{
						Id:         testDirId,
						ParentPath: testRootPath,
					}, nil)

				td.cRepo.EXPECT().
					Store(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content file store",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Select(gomock.Any(), gomock.Any()).
					Return(&model.Dir{
						Id:         testDirId,
						ParentPath: testRootPath,
					}, nil)

				td.cRepo.EXPECT().
					Store(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.cfRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

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

				td.cRepo.EXPECT().
					Store(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.cfRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil)

				td.fRepo.EXPECT().
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

			err := uc.Set(testCtx, tc.key, &content)

			require.ErrorIs(t, err, wantErr)
		})
	}
}
