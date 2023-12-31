package dir

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Select_Success(t *testing.T) {
	for _, tc := range []struct {
		name         string
		prepare      prepareFunc
		expFileCount uint64
	}{
		{
			name: "find dir and root",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: 1000,
						},
					}
					dirs = []model.Dir{{
						Id:         gofakeit.UUID(),
						FileCount:  10000,
						ParentPath: testRootPath,
					}, {
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return nil
			},
			expFileCount: testFileCount,
		},
		{
			name: "find dir and root with equal free and size",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: testSize,
						},
					}
					dirs = []model.Dir{{
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return nil
			},
			expFileCount: testFileCount,
		},
		{
			name: "find dir and root with fileCount + 1 = maxCount",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: 1000,
						},
					}
					dirs = []model.Dir{{
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return nil
			},
			expFileCount: testFileCount,
		},
		{
			name: "first dir with unknown root",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: 1000,
						},
					}
					dirs = []model.Dir{{
						Id:         gofakeit.UUID(),
						FileCount:  testFileCount,
						ParentPath: testRootPath2,
					}, {
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return nil
			},
			expFileCount: testFileCount,
		},
		{
			name: "first dir with to small root free space",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: 1000,
						},
						testRootPath2: {
							Free: testSize - 1,
						},
					}
					dirs = []model.Dir{{
						Id:         gofakeit.UUID(),
						FileCount:  testFileCount,
						ParentPath: testRootPath2,
					}, {
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return nil
			},
			expFileCount: testFileCount,
		},
		{
			name: "create new dir",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Path: testRootPath,
							Free: 1000,
						},
						testRootPath2: {
							Path: testRootPath2,
							Free: testSize - 1,
						},
					}
					dirs = []model.Dir{{
						Id:         gofakeit.UUID(),
						FileCount:  testFileCount,
						ParentPath: testRootPath2,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				td.dRepo.EXPECT().
					Create(gomock.Any(), model.Dir{
						Id:         testId,
						ParentPath: testRootPath,
					}).
					Return(nil)

				return nil
			},
			expFileCount: 0,
		},
		{
			name: "create new dir with empty dirs",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Path: testRootPath,
							Free: 1000,
						},
					}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(nil, nil)

				td.dRepo.EXPECT().
					Create(gomock.Any(), model.Dir{
						Id:         testId,
						ParentPath: testRootPath,
					}).
					Return(nil)

				return nil
			},
			expFileCount: 0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			_ = tc.prepare(td)

			uc := td.newUseCase()

			dir, err := uc.Select(context.Background(), testSize)

			require.NoError(t, err)
			require.Equal(t, &model.Dir{
				Id:         testId,
				FileCount:  tc.expFileCount,
				ParentPath: testRootPath,
			}, dir)
		})
	}
}

func TestUseCase_Select_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "no free space with empty rootMap",
			prepare: func(td *testDeps) error {
				var (
					dirs = []model.Dir{{
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(nil, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return fs_db.SizeErr
			},
		},
		{
			name: "no free space",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: testSize - 1,
						},
						testRootPath2: {
							Free: testSize - 1,
						},
					}
					dirs = []model.Dir{{
						Id:         testId,
						FileCount:  testFileCount,
						ParentPath: testRootPath,
					}}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(dirs, nil)

				return fs_db.SizeErr
			},
		},
		{
			name: "root get",
			prepare: func(td *testDeps) error {
				td.root.EXPECT().
					Get(gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "dir get",
			prepare: func(td *testDeps) error {
				td.root.EXPECT().
					Get(gomock.Any()).
					Return(nil, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "dir create",
			prepare: func(td *testDeps) error {
				var (
					rootMap = model.RootMap{
						testRootPath: {
							Free: testSize,
						},
					}
				)

				td.root.EXPECT().
					Get(gomock.Any()).
					Return(rootMap, nil)

				td.dRepo.EXPECT().
					Get(gomock.Any()).
					Return(nil, nil)

				td.dRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			dir, err := uc.Select(context.Background(), testSize)

			require.ErrorIs(t, err, wantErr)
			require.Nil(t, dir)
		})
	}
}
