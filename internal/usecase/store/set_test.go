package store

import (
	"path"
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
		dirs = model.Dirs{{
			Name: testDirName,
			Root: testRootPath,
			Free: testSize,
		}, {
			Name: testDirName2,
			Root: testRootPath,
			Free: testSize2,
		}, {
			Name: testDirName3,
			Root: testRootPath,
			Free: testSize3,
		}, {
			Name: testDirName4,
			Root: testRootPath,
			Free: testSize4,
		}}
		cFile = model.ContentFile{
			Id:     testContentId,
			Parent: path.Join(testRootPath, testDirName),
		}
		file = model.File{
			Key:       testKey,
			ContentId: testContentId,
		}
		content = testReader
		c       = testNewCloser(t, content, 2)
	)
	td := newTestDeps(t)

	td.dir.EXPECT().
		Get(gomock.Any()).
		Return(dirs, nil)

	td.cRepo.EXPECT().
		Store(gomock.Any(), path.Join(testRootPath, testDirName2, testContentId), gomock.Any()).
		Times(1).
		Return(model.NotEnoughSpaceError{
			Start: c,
		})

	td.cRepo.EXPECT().
		Store(gomock.Any(), path.Join(testRootPath, testDirName4, testContentId), gomock.Any()).
		Times(1).
		Return(model.NotEnoughSpaceError{
			Start: c,
		})

	td.cRepo.EXPECT().
		Store(gomock.Any(), path.Join(testRootPath, testDirName, testContentId), gomock.Any()).
		Times(1).
		Return(nil)

	td.cfRepo.EXPECT().
		Store(gomock.Any(), cFile).
		Return(nil)

	td.fRepo.EXPECT().
		Store(gomock.Any(), testTxId, file).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Set(testCtx, testKey, content)

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
			name: "size",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(nil, nil)

				return fs_db.SizeErr
			},
		},
		{
			name: "get dir",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "content store",
			key:  testKey,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(model.Dirs{{
						Name: testDirName,
						Root: testRootPath,
						Free: testSize,
					}}, nil)

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
					Get(gomock.Any()).
					Return(model.Dirs{{
						Name: testDirName,
						Root: testRootPath,
						Free: testSize,
					}}, nil)

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
					Get(gomock.Any()).
					Return(model.Dirs{{
						Name: testDirName,
						Root: testRootPath,
						Free: testSize,
					}}, nil)

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
				content = testReader
			)
			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			err := uc.Set(testCtx, tc.key, content)

			require.ErrorIs(t, err, wantErr)
		})
	}
}
